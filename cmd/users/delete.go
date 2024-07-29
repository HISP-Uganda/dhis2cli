package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	DeleteCmd.Flags().StringVarP(&userID, "uid", "", "", "UID for user to delete.\nNote: This is required if 'ids-file' not provided")
	DeleteCmd.Flags().StringVarP(&idsFile, "ids-file", "", "", "File with UIDs (one per line) for users to delete. \nNote: required if 'id' flag not provided")
	DeleteCmd.MarkFlagsOneRequired("uid", "ids-file")
	DeleteCmd.MarkFlagsMutuallyExclusive("uid", "ids-file")
}

// DeleteCmd is the command for deleting users
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting user...")
	},
}
