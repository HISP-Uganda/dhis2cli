package sms

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var SentSMSCmd = &cobra.Command{
	Use:   "sent",
	Short: "Sent SMS messages",
	Run: func(cmd *cobra.Command, args []string) {

		defaultParams := map[string]any{
			"fields": "id,message,status,date,recipients",
			"order":  "date:desc",
		}
		additionalParams := map[string]any{}
		if status != "" {
			additionalParams["filter"] = fmt.Sprintf("status:eq:%s", status)
		}
		excludeKeys := []string{"query", "paging"}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, excludeKeys)
		fmt.Println("Params:", params)
		utils.FetchResourceAndDisplay(client.Dhis2Client, "/sms/outbound", params, "outboundsmss", config.OutputFormat)
		//resp, err := client.Dhis2Client.GetResource(
		//	"/sms/outbound", params)
		//if err != nil {
		//	fmt.Printf("Error fetching sent SMS: %v\n", err)
		//	return
		//}
		//fmt.Println(string(resp.Body()))
		//err = json.Unmarshal(resp.Body(), &resp)
		//if err != nil {
		//	fmt.Printf("Error unmarshalling users list: %v\n", err)
		//	return
		//}
		//err = utils.DisplayTable(resp)
		//// err = utils.DisplayTable2(users["users"], true)
		//if err != nil {
		//	fmt.Println("Error:", err)
		//}

	},
}

func init() {
	SentSMSCmd.Flags().StringVarP(&status, "status", "", "", "The status of the messages to show")
}
