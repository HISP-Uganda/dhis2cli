package apps

import "github.com/spf13/cobra"

func init() {
	AppsCmd.AddCommand(ListAppsCmd)
	AppsCmd.AddCommand(ManualInstallCmd)
	AppsCmd.AddCommand(DeleteAppCmd)
}

var AppsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Manage apps",
}
