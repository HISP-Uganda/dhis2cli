package apps

import "github.com/spf13/cobra"

func init() {
	AppsCmd.AddCommand(ListAppsCmd)
}

var AppsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Manage apps",
}
