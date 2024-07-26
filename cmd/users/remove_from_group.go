package users

import "github.com/spf13/cobra"

var RemoveFromGroupCmd = &cobra.Command{
	Use:   "removeFromGroup",
	Short: "Remove user(s) from a group",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement user removal from a group
	},
}
