package users

import "github.com/spf13/cobra"

var RemoveFromOrgUnitCmd = &cobra.Command{
	Use:   "removeFromOrgUnit",
	Short: "Remove user(s) from an organization unit",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement user removal from an organization unit
	},
}
