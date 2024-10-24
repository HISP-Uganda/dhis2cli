package maintenance

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var limit int

func init() {
	IdentifiersCmd.Flags().IntVar(&limit, "limit", 1, "How many identifiers you want to be returned")
}

var IdentifiersCmd = &cobra.Command{
	Use:   "identifiers",
	Short: "Generate identifiers",
	Run: func(cmd *cobra.Command, args []string) {
		params := map[string]any{
			"limit": limit,
		}
		utils.FetchResourceAndDisplay2(client.Dhis2Client, "system/id", params, "", "json")
	},
}
