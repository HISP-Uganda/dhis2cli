package geojson

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	GeoJsonCmd.AddCommand(ExportCmd)
	// GeoJsonCmd.AddCommand(ImportCmd)
	ExportCmd.Flags().StringVarP(&exportFile, "output-file", "", "Organisation_units.geojson", "The output file")
	ExportCmd.Flags().IntVarP(&ouLevel, "ou-level", "", 1, "The ou level")
}

var GeoJsonCmd = &cobra.Command{
	Use:   "geojson",
	Short: "Export or Import  GeoJson",
}
var exportFile string
var ouLevel int

var ExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export GeoJson",
	Run: func(cmd *cobra.Command, args []string) {
		defaultParams := map[string]any{
			"fields": "id,name,geometry",
			"paging": "false",
		}
		additionalParams := map[string]any{
			"level": ouLevel,
		}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		resp, err := client.Dhis2Client.GetResource("/organisationUnits.geojson", params)
		if err != nil {
			fmt.Printf("Error fetching GeoJson: %v\n", err)
			return
		}

		if err := os.WriteFile(exportFile, resp.Body(), 0644); err != nil {
			fmt.Printf("Failed writing file: %v", err)
		}
		fmt.Printf("✅ GeoJSON written to %s\n", exportFile)
	},
}
