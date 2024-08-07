package trackedentitytypes

import "github.com/spf13/cobra"

func init() {
	TrackedEntityTypeCmd.AddCommand(ListCmd)
}

var program string
var orgUnit string
var TrackedEntityTypeCmd = &cobra.Command{
	Use:   "trackedEntityType",
	Short: "Manage tracked entity type",
}
