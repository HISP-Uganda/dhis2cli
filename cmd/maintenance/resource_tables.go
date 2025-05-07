package maintenance

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var ResourceTablesCmd = &cobra.Command{
	Use:   "resourceTables",
	Short: "Generate resource tables",
	Run: func(cmd *cobra.Command, args []string) {
		params := make(map[string]interface{})
		utils.PostResourceAndDisplay(client.Dhis2Client, "resourceTables", params, nil, "", "json")
	},
}
