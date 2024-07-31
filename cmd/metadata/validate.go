package metadata

import (
	"fmt"
	"github.com/spf13/cobra"
)

var payload string
var payloadFile string

// ValidateCmd is a command to validate payloads
var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate metadata payloads",
	Run: func(cmd *cobra.Command, args []string) {
		// Validate metadata objects logic here
		// use the /api/schemas/resource endpoint
		fmt.Printf("This command is %s", cmd.Name())
	},
}

func init() {
	ValidateCmd.Flags().StringVarP(&payload, "payload", "", "", "The payload to validate")
	ValidateCmd.Flags().StringVarP(&payloadFile, "payload-file", "", "", "The payload to validate")
	// ValidateCmd.MarkFlagsOneRequired("payload", "payload-file")
	ValidateCmd.MarkFlagsMutuallyExclusive("payload", "payload-file")
	// ValidateCmd.MarkFlagCustom("payload-file", "payload-file")
	_ = ValidateCmd.MarkFlagFilename("payload-file", ".json")
}
