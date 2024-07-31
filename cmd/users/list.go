package users

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var format = "table"

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		defaultParams := map[string]any{
			"fields":          "id,displayName,access,email,username,lastLogin",
			"order":           "firstName:asc,surname:asc",
			"paging":          "true",
			"userOrgUnits":    "true",
			"includeChildren": "true",
			"selfRegistered":  "false",
		}
		additionalParams := map[string]any{}
		// excludeKeys := []string{}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		utils.FetchResourceAndDisplay2(client.Dhis2Client, "/users", params, "users", config.OutputFormat)

	},
}
