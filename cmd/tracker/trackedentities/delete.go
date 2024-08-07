package trackedentities

import (
	"dhis2cli/client"
	"dhis2cli/models/tracker"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var teID string
var idsFile string
var validate bool

func init() {
	DeleteCmd.Flags().StringVar(&teID, "uid", "", "UID for tracked entity to delete.\nNote: This is required if 'ids-file' not provided")
	DeleteCmd.Flags().StringVar(&idsFile, "ids-file", "", "File with UIDs (one per line) for tracked entities to delete. \nNote: required if 'id' flag not provided")
	DeleteCmd.Flags().BoolVar(&validate, "validate", false, "Validate tracked entities deletion")
	DeleteCmd.MarkFlagsOneRequired("uid", "ids-file")
	DeleteCmd.MarkFlagsMutuallyExclusive("uid", "ids-file")
	_ = DeleteCmd.MarkFlagFilename("ids-file", "csv", "txt", ".csv", ".txt")
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete tracked entities",
	Run: func(cmd *cobra.Command, args []string) {
		params := map[string]interface{}{
			"importStrategy": "DELETE",
			"importMode":     "COMMIT",
		}
		if validate {
			params["importMode"] = "VALIDATE"
		}
		var payload tracker.FlatPayload
		if teID != "" {
			// fmt.Printf("Deleting tracked entity with ID: %s\n", teID)
			trackedEntities := []tracker.TrackedEntity{
				tracker.TrackedEntity{
					TrackedEntity: teID,
				},
			}
			payload = tracker.FlatPayload{
				TrackedEntities: trackedEntities,
			}

		} else if idsFile != "" {
			// fmt.Println("Deleting tracked entities from file")
			lines, _ := utils.ReadCSV(idsFile)
			var trackedEntities []tracker.TrackedEntity
			for _, line := range lines {
				// fmt.Printf("Deleting tracked entity with ID: %s\n", line)
				// Add tracked entity to trackedEntities slice
				if len(line) > 0 {
					trackedEntities = append(trackedEntities, tracker.TrackedEntity{
						TrackedEntity: line[0],
					})

				}
				// Delete tracked entity with ID from here
			}
			payload = tracker.FlatPayload{
				TrackedEntities: trackedEntities,
			}
		}
		utils.PostResourceAndDisplay(client.Dhis2Client, "tracker", params, payload, "", "json")
	},
}
