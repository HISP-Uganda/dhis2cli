package sms

import (
	"dhis2cli/client"
	"dhis2cli/models"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var recipients, groups, orgunits string
var SendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send SMS",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement sending SMS message
	},
}

func init() {
	// SendCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "SMS message")
	SendCmd.AddCommand(toPhoneNumbersCmd)

	toPhoneNumbersCmd.Flags().StringVarP(&message, "message", "m", "", "SMS message")
	toPhoneNumbersCmd.Flags().StringVarP(&messageFile, "message-file", "M", "", "SMS message file")
	toPhoneNumbersCmd.Flags().StringVarP(&recipients, "recipients", "t", "", "Message recipients. Comma-separated if multiple")
	_ = toPhoneNumbersCmd.MarkFlagRequired("recipients")
	toPhoneNumbersCmd.MarkFlagsOneRequired("message", "message-file")

	SendCmd.AddCommand(toGroupsCmd)
	toGroupsCmd.Flags().StringVarP(&groups, "groups", "g", "", "Recipient groups. Comma-separated if multiple")

	SendCmd.AddCommand(toUserAssignedToOrgunitCmd)
	toUserAssignedToOrgunitCmd.Flags().StringVarP(&orgunits, "orgunits", "o", "", "Organization unit uid: Comma-separated if multiple")

	SendCmd.AddCommand(toOrgunitPhoneNumberCmd)
	toOrgunitPhoneNumberCmd.Flags().StringVarP(&orgunits, "orgunits", "o", "", "Organization unit uid: Comma-separated if multiple")
}

var toPhoneNumbersCmd = &cobra.Command{
	Use:   "toPhoneNumbers",
	Short: "Send SMS to phone numbers",
	Example: `
dhis2 sms send toPhoneNumbers -t "256753475676" -m "Test message"

dhis2 sms send toPhoneNumbers --message "Hello, world!" --recipients "+256782820208,+256753475676"

dhis2 sms send toPhoneNumbers --message-file "message.txt" --recipients "+1234567890,+9876543210"
`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement sending SMS to phone numbers
		recipients := strings.Split(recipients, ",")
		uniqueNumbers := utils.RemoveDuplicates(recipients)
		if messageFile != "" {
			content, err := utils.ReadFile(messageFile)
			if err != nil {
				fmt.Printf("Error reading message file: %v\n", err)
				return
			}
			message = content
		}
		if message == "" {
			fmt.Println("Message cannot be empty")
			return
		}
		payload := models.SMSPayload{
			Message:    message,
			Recipients: uniqueNumbers,
		}
		resp, err := client.Dhis2Client.PostResource(
			"/sms/outbound", nil, payload)
		if err != nil {
			fmt.Printf("Error sending SMS: %v\n", err)
			return
		}
		fmt.Println(string(resp.Body()))

	},
}

var toGroupsCmd = &cobra.Command{
	Use:   "toGroups",
	Short: "Send SMS to groups",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement sending SMS to groups
	},
}

var toUserAssignedToOrgunitCmd = &cobra.Command{
	Use:   "toUserAssignedToOrgunit",
	Short: "Send SMS to users assigned to an organization unit",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement sending SMS to users assigned to an organization unit
	},
}

var toOrgunitPhoneNumberCmd = &cobra.Command{
	Use:   "toOrgunitPhoneNumber",
	Short: "Send SMS to an organization unit phone number",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement sending SMS to an organization unit phone number
	},
}
