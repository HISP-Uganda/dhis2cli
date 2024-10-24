package datastore

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var data string
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing data Store namespace key",
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "dataStore"
		if namespace != "" {
			endpoint = fmt.Sprintf("%s/%s", endpoint, namespace)
		}
		if key != "" && namespace != "" {
			endpoint = fmt.Sprintf("%s/%s", endpoint, key)
		}
		resp, err := client.Dhis2Client.PutResource(endpoint, data)
		if err != nil {
			// Log failed to update
			fmt.Printf("Failed to update data store: %v\n", err)
			return
		}
		fmt.Printf(utils.PrintResponse(string(resp.Body()), true))
	},
}

func init() {
	UpdateCmd.Flags().StringVar(&namespace, "namespace", "", "Namespace for the data store")
	UpdateCmd.Flags().StringVar(&key, "key", "", "key to retrieve from namespace")
	UpdateCmd.Flags().StringVar(&data, "data", "", "JSON representation of the key's value")
}
