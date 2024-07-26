package users

import "github.com/spf13/cobra"

var AddToOrgUnitCmd = &cobra.Command{
	Use:   "addToOrgUnit",
	Short: "Add user(s) to an organization unit",
	Run: func(cmd *cobra.Command, args []string) {
		// Add user to org unit logic here
	},
}
