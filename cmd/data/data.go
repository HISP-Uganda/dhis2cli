package data

import "github.com/spf13/cobra"

var DataCmd = &cobra.Command{
	Use:   "data",
	Short: "Sending and receiving data values",
}

func init() {
	DataCmd.AddCommand(DataValueSetTemplateCmd)
	DataCmd.AddCommand(SendDataValuesCmd)
}
