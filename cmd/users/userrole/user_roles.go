package userrole

import "github.com/spf13/cobra"

func init() {

}

var userRolesCmd = &cobra.Command{
	Use:   "userRoles",
	Short: "Manage user roles",
	//Run: func(cmd *cobra.Command, args []string) {
	//    fmt.Println("User roles management...")
	//},
}
