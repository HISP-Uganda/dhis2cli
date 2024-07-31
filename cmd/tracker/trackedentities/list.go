package trackedentities

import "github.com/spf13/cobra"

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tracked entities",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement fetching and displaying tracked entities
	},
}
