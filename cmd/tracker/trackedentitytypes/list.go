package trackedentitytypes

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tracked entity types",
	Run: func(cmd *cobra.Command, args []string) {
		defaultParams := map[string]any{
			"fields":   "id,displayName,shortName,lastUpdated,created",
			"filters":  []string{"name:ne:default"},
			"order":    "displayName:ASC",
			"paging":   "true",
			"page":     "1",
			"pageSize": "10",
		}
		additionalParams := map[string]any{}
		// excludeKeys := []string{}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		utils.FetchResourceAndDisplay2(client.Dhis2Client, "/trackedEntityTypes", params, "trackedEntityTypes", config.OutputFormat)
	},
}
