package cmd

import (
	"dhis2cli/client"
	"dhis2cli/cmd/apps"
	"dhis2cli/cmd/sms"
	"dhis2cli/cmd/tracker"
	"dhis2cli/cmd/users"
	"dhis2cli/config"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var cfgFile string

var version = "1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: version,
	Use:     "dhis2",
	Short:   "DHIS2 on your command-line",
	Long: `A DHIS2 commandline application for the common DHIS2 tasks:
The DHIS2 CLI brings some of DHIS2's power to your command-line.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cfgFile != "" {
			config.LoadConfig(cfgFile)
		} else {
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)
			path := filepath.Join(home, ".dhis2cli.yaml")
			config.LoadConfig(path)
			// config.LoadConfig("config.yaml")
			client.InitServer()
			client.Dhis2Client, _ = client.Dhis2Server.NewClient()
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(
		&cfgFile, "config", "c", "", "config file (default is $HOME/.dhis2cli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "V", false, "Show some debug information")
	rootCmd.PersistentFlags().StringVarP(&config.GlobalParams.Paging, "paging", "", "true", "Whether to return lists of elements in pages.")
	rootCmd.PersistentFlags().IntVarP(&config.GlobalParams.Page, "page", "p", 1, "Page number to return.")
	rootCmd.PersistentFlags().IntVarP(&config.GlobalParams.PageSize, "page-size", "P", 10, "Number of elements to return for each page.")
	rootCmd.PersistentFlags().StringVarP(&config.GlobalParams.Fields, "fields", "f", "", "Fields to display")
	rootCmd.PersistentFlags().StringSliceVarP(&config.GlobalParams.Filter, "filters", "F", []string{}, "Filters to apply")
	rootCmd.PersistentFlags().StringVarP(&config.GlobalParams.Order, "order", "O", "", "How to order the output:\nproperty:asc/iasc/desc/idesc")
	rootCmd.PersistentFlags().StringVarP(&config.GlobalParams.Query, "query", "q", "", "Query term used to search through all fields")
	rootCmd.PersistentFlags().StringVarP(&config.OutputFormat, "format", "", "table", "Output format: table/json/csv")
	rootCmd.PersistentFlags().StringVarP(&config.OutputFile, "output-file", "o", "", "Output file")
	rootCmd.PersistentFlags().BoolP("indent", "i", false, "Whether to indent JSON output")

	rootCmd.SetVersionTemplate(fmt.Sprintf("DHIS2 CLI: %s\n", version))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(users.UsersCmd)
	rootCmd.AddCommand(sms.SmsCmd)
	rootCmd.AddCommand(apps.AppsCmd)
	rootCmd.AddCommand(tracker.TrackerCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".dhis2cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".dhis2cli")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("DHIS2CLI")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
