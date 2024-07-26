package users

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ImportUsersCmd = &cobra.Command{
	Use:   "import",
	Short: "Import users from a file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Importing users from CSV file...")
	},
}
