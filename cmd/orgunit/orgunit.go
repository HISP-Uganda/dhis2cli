package orgunit

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	OrgUnitCmd.AddCommand(ListCmd)
}

var OrgUnitCmd = &cobra.Command{
	Use:   "orgUnit",
	Short: "Manage organization units",
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all organization units",
	Run: func(cmd *cobra.Command, args []string) {
		defaultParams := map[string]any{
			"fields":   "id,displayName,level,path,lastUpdated",
			"filters":  []string{"name:ne:default"},
			"paging":   "true",
			"page":     "1",
			"pageSize": "10",
		}
		additionalParams := map[string]any{}
		// excludeKeys := []string{}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		utils.FetchResourceAndDisplay2(client.Dhis2Client, "/organisationUnits", params, "organisationUnits", config.OutputFormat)
	},
}
