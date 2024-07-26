package users

import (
	"github.com/spf13/cobra"
)

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
}
