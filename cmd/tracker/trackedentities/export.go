package trackedentities

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var exportFormat string
var compressionType string
var ExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export tracked entities",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement exporting tracked entities
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
		mimeType, _ := utils.GetContentType(exportFormat, compressionType)
		fileName := fmt.Sprintf("trackedEntities.%s", exportFormat)
		switch compressionType {
		case "zip":
			fileName = fmt.Sprintf("%s.zip", fileName)
		case "gzip":
			fileName = fmt.Sprintf("%s.gz", fileName)
		}

		utils.FetchExport(client.Dhis2Client, "tracker/trackedEntities", params, mimeType, fileName)

	},
}

func init() {
	ExportCmd.Flags().StringVar(&exportFormat, "exportFormat", "json", "The export format")
	ExportCmd.Flags().StringVar(&compressionType, "compression", "none", "The compression type. zip or gzip")
}
