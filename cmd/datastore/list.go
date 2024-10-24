package datastore

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var namespace string
var key string
var metaData bool
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List data stores",
	Run: func(cmd *cobra.Command, args []string) {
		// List resources logic here
		defaultParams := map[string]any{
			"fields": "",
		}
		additionalParams := map[string]any{}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		endpoint := "dataStore"
		if namespace != "" {
			endpoint = fmt.Sprintf("%s/%s", endpoint, namespace)
		}
		if key != "" && namespace != "" {
			endpoint = fmt.Sprintf("%s/%s", endpoint, key)
		}
		if key != "" && metaData {
			endpoint = fmt.Sprintf("%s/metaData", endpoint)
		}

		utils.FetchResourceAndDisplay2(client.Dhis2Client, endpoint, params, "", "json")
	},
}

func init() {
	ListCmd.Flags().StringVar(&key, "key", "", "key to retrieve from namespace")
	ListCmd.Flags().BoolVar(&metaData, "meta-data", false, "Whether to retrieve meta-data for an existing key from a namespace")
	_ = ListCmd.MarkFlagRequired("namespace")
	// ListCmd.MarkFlagsRequiredTogether("key")
}
