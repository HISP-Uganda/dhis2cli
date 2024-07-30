package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	AddToGroupCmd.Flags().StringVarP(&groupId, "user-group", "g", "", "UID for user group.")
	AddToGroupCmd.Flags().StringVarP(&userID, "user", "u", "", "UID for user to add to group.\nNote: This is required if 'user-ids-file' not provided")
	AddToGroupCmd.Flags().StringVarP(&userIdsFile, "user-ids-file", "", "", "File with user UIDs (one per line) for users to assign to user role. \nNote: required if 'id' flag not provided")
	AddToGroupCmd.MarkFlagsOneRequired("user", "user-ids-file")
	AddToGroupCmd.MarkFlagsMutuallyExclusive("user", "user-ids-file")
	var _ = AssignRoleCmd.MarkFlagRequired("user-role")
}

var AddToGroupCmd = &cobra.Command{
	Use:   "addToGroup",
	Short: "Add user(s) to a group",
	Run: func(cmd *cobra.Command, args []string) {
		// Add user to group logic here
		fmt.Printf("Add user: %s to Group: %s", userID, groupId)
	},
}
