package apps

import (
	"dhis2cli/client"
	"fmt"
	"github.com/spf13/cobra"
)

var appFile string
var ManualInstallCmd = &cobra.Command{
	Use:   "manualInstall",
	Short: "Manually install an app",
	Run: func(cmd *cobra.Command, args []string) {
		// Print installing app with its details
		fmt.Printf("Installing app: %s\n", appFile)

		_, err := client.Dhis2Client.PostFileResource("/apps", "file", appFile)
		if err != nil {
			// Log failed to install
			fmt.Printf("Failed to install app: %v\n", err)
			return
		}
		fmt.Println("App installed successfully:")
	},
}

func init() {
	ManualInstallCmd.Flags().StringVarP(&appFile, "app-file", "", "", "Path to app zip build file")
	_ = ManualInstallCmd.MarkFlagRequired("app-file")
}
