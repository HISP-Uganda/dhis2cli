package datastore

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing data Store namespace/keys",
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "dataStore"
		if namespace != "" {
			endpoint = fmt.Sprintf("%s/%s", endpoint, namespace)
		}
		if key != "" && namespace != "" {
			endpoint = fmt.Sprintf("%s/%s", endpoint, key)
		}
		resp, err := client.Dhis2Client.DeleteResource(endpoint)
		if err != nil {
			// Log failed to delete
			fmt.Printf("Failed to delete data store: %v\n", err)
			return
		}
		fmt.Printf(utils.PrintResponse(string(resp.Body()), true))
	},
}

func init() {
	DeleteCmd.Flags().StringVar(&namespace, "namespace", "", "Namespace for the data store")
	DeleteCmd.Flags().StringVar(&key, "key", "", "key to retrieve from namespace")
}
