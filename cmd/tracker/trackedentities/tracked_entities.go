package trackedentities

import "github.com/spf13/cobra"

var TrackedEntities = &cobra.Command{
	Use:   "trackedEntities",
	Short: "Manage tracked entities",
}

func init() {
	TrackedEntities.AddCommand(ListCmd)
	TrackedEntities.AddCommand(DeleteCmd)
}
