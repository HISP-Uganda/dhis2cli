package metadata

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var ViewResourcesCmd = &cobra.Command{
	Use:   "viewResources",
	Short: "View metadata resources",
	Run: func(cmd *cobra.Command, args []string) {
		// View metadata resource logic here
		utils.FetchResourceAndDisplay2(client.Dhis2Client, "/resources", nil, "resources", config.OutputFormat)
	},
}
