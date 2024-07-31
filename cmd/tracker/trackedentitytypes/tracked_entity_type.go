package trackedentitytypes

import "github.com/spf13/cobra"

func init() {
	TrackedEntityTypeCmd.AddCommand(ListCmd)
}

var TrackedEntityTypeCmd = &cobra.Command{
	Use:   "trackedEntityType",
	Short: "Manage tracked entity type",
}
