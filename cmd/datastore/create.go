package datastore

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new data Store namespace/key",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new resource logic here
		endpoint := "dataStore"
		if namespace != "" {
			endpoint = fmt.Sprintf("%s/%s", endpoint, namespace)
		}
		if key != "" && namespace != "" {
			endpoint = fmt.Sprintf("%s/%s", endpoint, key)
		}
		resp, err := client.Dhis2Client.PostResource(endpoint, nil, data)
		if err != nil {
			// Log failed to update
			fmt.Printf("Failed to create new key and value for a namespace: %v\n", err)
			return
		}
		fmt.Printf(utils.PrintResponse(string(resp.Body()), true))
	},
}
