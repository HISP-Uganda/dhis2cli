package users

import (
	"dhis2cli/cmd/users/usergroup"
	"dhis2cli/cmd/users/userrole"
	"github.com/spf13/cobra"
)

var userID string
var idsFile string
var groupId string

var UsersCmd = &cobra.Command{
	Use:   "users",
	Short: "Manage users",
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("Users management...")
	//},
}

func init() {
	UsersCmd.AddCommand(ListCmd)
	UsersCmd.AddCommand(ImportUsersCmd)
	UsersCmd.AddCommand(DeleteCmd)
	UsersCmd.AddCommand(AddToGroupCmd)
	UsersCmd.AddCommand(RemoveFromGroupCmd)
	UsersCmd.AddCommand(AssignRoleCmd)
	UsersCmd.AddCommand(RemoveRoleCmd)
	UsersCmd.AddCommand(AddToOrgUnitCmd)
	UsersCmd.AddCommand(RemoveFromOrgUnitCmd)
	UsersCmd.AddCommand(SetPasswordCmd)
	UsersCmd.AddCommand(usergroup.UserGroupsCmd)
	UsersCmd.AddCommand(userrole.UserRolesCmd)
}
