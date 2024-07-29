package sms

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

var ReceivedSMSCmd = &cobra.Command{
	Use:   "received",
	Short: "Receive SMS messages",
	Run: func(cmd *cobra.Command, args []string) {
		defaultParams := map[string]string{
			"fields": "id,text,originator,smsstatus,user[userCredentials[username]],receiveddate",
			"order":  "receiveddate:desc",
		}
		additionalParams := map[string]string{}
		if status != "" {
			additionalParams["filter"] = fmt.Sprintf("status:eq:%s", status)
		}
		excludeKeys := []string{"query", "paging"}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, excludeKeys)
		fmt.Println("Params:", params)
		resp, err := client.Dhis2Client.GetResource(
			"/sms/inbound", params)
		if err != nil {
			fmt.Printf("Error fetching sent SMS: %v\n", err)
			return
		}
		fmt.Println(string(resp.Body()))
		err = json.Unmarshal(resp.Body(), &resp)
		if err != nil {
			fmt.Printf("Error unmarshalling users list: %v\n", err)
			return
		}
		err = utils.DisplayTable(resp)
		// err = utils.DisplayTable2(users["users"], true)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	ReceivedSMSCmd.Flags().StringVarP(&status, "status", "", "", "The status of the messages to show")
}
