package usergroup

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	UserGroupsCmd.AddCommand(ListUserGroupsCmd)
}

var UserGroupsCmd = &cobra.Command{
	Use:   "userGroups",
	Short: "Manage user groups",
	//Run: func(cmd *cobra.Command, args []string) {
	//    fmt.Println("User groups management...")
	//},
}

var ListUserGroupsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all user groups",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement listing all user groups
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
		fmt.Println("PARAMS", params)
		utils.FetchResourceAndDisplay(client.Dhis2Client, "/userGroups", params, "userGroups", config.OutputFormat)

	},
}
