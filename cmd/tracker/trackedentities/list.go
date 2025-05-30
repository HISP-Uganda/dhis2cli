package trackedentities

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tracked entities",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement fetching and displaying tracked entities
		defaultParams := map[string]any{
			"fields":     "*",
			"order":      "createdAt:desc",
			"paging":     "true",
			"page":       "1",
			"pageSize":   "10",
			"totalPages": config.GlobalParams.TotalPages,
		}
		additionalParams := map[string]any{}
		program, _ := cmd.Flags().GetString("program")
		if program != "" {
			additionalParams["program"] = program
		}
		teType, _ := cmd.Flags().GetString("trackedEntityType")
		if teType != "" {
			additionalParams["trackedEntityType"] = teType
		}
		orgUnitMode, _ := cmd.Flags().GetString("orgunitMode")
		if orgUnitMode != "" {
			additionalParams["ouMode"] = orgUnitMode
			additionalParams["orgunitMode"] = orgUnitMode
		}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		resource := ""
		switch config.OutputFormat {
		case "csv", "table":
			resource = "instances"
		}

		utils.FetchResourceAndDisplay2(client.Dhis2Client, "tracker/trackedEntities", params, resource, config.OutputFormat)

	},
}
