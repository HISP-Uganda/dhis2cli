package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var userIdsFile string
var roleID string

func init() {
	AssignRoleCmd.Flags().StringVarP(&roleID, "user-role", "r", "", "UID for user role.")
	AssignRoleCmd.Flags().StringVarP(&userID, "user", "u", "", "UID for user to assign role.\nNote: This is required if 'user-ids-file' not provided")
	AssignRoleCmd.Flags().StringVarP(&userIdsFile, "user-ids-file", "f", "", "File with user UIDs (one per line) for users to assign to user role. \nNote: required if 'id' flag not provided")
	AssignRoleCmd.MarkFlagsOneRequired("user", "user-ids-file")
	AssignRoleCmd.MarkFlagsMutuallyExclusive("user", "user-ids-file")
	var _ = AssignRoleCmd.MarkFlagRequired("user-role")
}

var AssignRoleCmd = &cobra.Command{
	Use:   "assignRole",
	Short: "Assign a role to user(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement role assignment logic
		fmt.Printf("Assign role: %s to User: %s\n", roleID, userID)
	},
}
