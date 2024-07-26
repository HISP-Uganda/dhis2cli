package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var userID string
var idsFile string

func init() {
	DeleteCmd.Flags().StringVarP(&userID, "id", "i", "", "UID for user to delete.\nNote: This is required if 'ids-file' not provided")
	DeleteCmd.Flags().StringVarP(&idsFile, "ids-file", "f", "", "File with UIDs (one per line) for users to delete. \nNote: required if 'id' flag not provided")
	DeleteCmd.MarkFlagsOneRequired("id", "ids-file")
	DeleteCmd.MarkFlagsMutuallyExclusive("id", "ids-file")
}

// DeleteCmd is the command for deleting users
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting user...")
	},
}
