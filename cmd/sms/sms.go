package sms

import "github.com/spf13/cobra"

var message, messageFile string
var status string

var SmsCmd = &cobra.Command{
	Use:   "sms",
	Short: "SMS Command",
	// Run: func(cmd *cobra.Command, args []string) {
	//    fmt.Println("Sending SMS messages...")
	//},
}

func init() {
	SmsCmd.AddCommand(SendCmd)
	SmsCmd.AddCommand(SentSMSCmd)
	SmsCmd.AddCommand(ReceivedSMSCmd)

	//SmsCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "SMS message")
	//SmsCmd.PersistentFlags().StringVarP(&messageFile, "message-file", "M", "", "SMS message file")
}
