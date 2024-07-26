package users

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

}

// RemoveRoleCmd represents the removeRole command
var RemoveRoleCmd = &cobra.Command{
	Use:   "removeRole",
	Short: "Remove user(s) from a role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("removeRole called")
	},
}
