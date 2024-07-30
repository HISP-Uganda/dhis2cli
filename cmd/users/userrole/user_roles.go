package userrole

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	UserRolesCmd.AddCommand(ListCmd)
}

var UserRolesCmd = &cobra.Command{
	Use:   "userRoles",
	Short: "Manage user roles",
	//Run: func(cmd *cobra.Command, args []string) {
	//    fmt.Println("User roles management...")
	//},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all user roles",
	Run: func(cmd *cobra.Command, args []string) {
		defaultParams := map[string]any{
			"fields":   "id,displayName,user[id,displayName]",
			"order":    "name:asc",
			"paging":   "true",
			"page":     "1",
			"pageSize": "50",
		}
		additionalParams := map[string]any{}
		// excludeKeys := []string{}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		utils.FetchResourceAndDisplay(client.Dhis2Client, "/userRoles", params, "userRoles", config.OutputFormat)
	},
}
