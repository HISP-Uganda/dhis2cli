package maintenance

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check DHIS2 username and password combination is correct",
	Run: func(cmd *cobra.Command, args []string) {
		// Ping the DHIS2 instance
		utils.FetchResourceAndDisplay2(client.Dhis2Client, "system/ping", nil, "", "string")
	},
}

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display DHIS2 instance information",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch and display DHIS2 instance information
		utils.FetchResourceAndDisplay2(client.Dhis2Client, "system/info", nil, "", "json")
	},
}
