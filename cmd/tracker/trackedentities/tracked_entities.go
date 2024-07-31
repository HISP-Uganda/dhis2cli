package trackedentities

import "github.com/spf13/cobra"

var TrackedEntitiesCmd = &cobra.Command{
	Use:   "trackedEntities",
	Short: "Manage tracked entities",
}

func init() {
	TrackedEntitiesCmd.AddCommand(ListCmd)
	TrackedEntitiesCmd.AddCommand(DeleteCmd)
}
