package users

import "github.com/spf13/cobra"

func init() {
	RemoveFromGroupCmd.Flags().StringVarP(&groupId, "user-group", "", "", "UID for user group.")
	RemoveFromGroupCmd.Flags().StringVarP(&userID, "user", "", "", "UID for user to add to group.\nNote: This is required if 'user-ids-file' not provided")
	RemoveFromGroupCmd.Flags().StringVarP(&userIdsFile, "user-ids-file", "", "", "File with user UIDs (one per line) for users to assign to user role. \nNote: required if 'id' flag not provided")
	RemoveFromGroupCmd.MarkFlagsOneRequired("user", "user-ids-file")
	RemoveFromGroupCmd.MarkFlagsMutuallyExclusive("user", "user-ids-file")
	var _ = AssignRoleCmd.MarkFlagRequired("user-role")
}

var RemoveFromGroupCmd = &cobra.Command{
	Use:   "removeFromGroup",
	Short: "Remove user(s) from a group",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement user removal from a group
	},
}
