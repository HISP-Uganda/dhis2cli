package trackedentities

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var payload string
var payloadFile string
var streamingOn bool

var ImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import tracked entities from a file",
	Run: func(cmd *cobra.Command, args []string) {
		resourcePath := "tracker"
		var toValidate string
		defaultParams := map[string]any{
			"importStrategy": "CREATE_AND_UPDATE",
			"importMode":     "COMMIT",
			"atomicMode":     "ALL",
			"async":          "true",
		}
		additionalParams := map[string]any{}
		// excludeKeys := []string{}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		if payloadFile != "" {
			// Read contents from the specified payload file
			var err error
			if streamingOn {
				// Stream file while reading off TrackedEntities
				return
			} else {
				toValidate, err = utils.ReadFileToString(payloadFile)
				if err != nil {
					fmt.Printf("Error reading payload file: %v\n", err)
					return
				}
			}
		} else if payload != "" {
			// Use the provided payload string
			toValidate = payload
		} else {
			fmt.Println("No payload provided.")
			return
		}
		fmt.Printf("%v\n", toValidate)
		utils.PostResourceAndDisplay(client.Dhis2Client, resourcePath, params, toValidate, "", "json")
	},
}

func init() {
	ImportCmd.Flags().StringVar(&payload, "payload", "", "The payload to validate")
	ImportCmd.Flags().StringVar(&payloadFile, "payload-file", "", "The file with payload to validate")
	ImportCmd.Flags().BoolVar(&streamingOn, "streaming", false, "Turn on streaming")
	ImportCmd.MarkFlagsMutuallyExclusive("payload", "payload-file")
	ImportCmd.MarkFlagsOneRequired("payload-file", "payload")
	_ = ImportCmd.MarkFlagFilename("payload-file", ".json")
}
