package apps

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var ListAppsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all apps",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement listing all apps
		defaultParams := map[string]any{
			"fields": "key,name,version,description",
		}
		additionalParams := map[string]any{}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		utils.FetchResourceAndDisplay2(client.Dhis2Client, "/apps", params, "", config.OutputFormat)
	},
}
