package trackedentities

import "github.com/spf13/cobra"

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete tracked entities",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement deleting a tracked entity
	},
}
