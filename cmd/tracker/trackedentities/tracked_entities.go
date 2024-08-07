package trackedentities

import (
	"github.com/spf13/cobra"
)

var program string
var orgUnit string
var orgUnitMode string
var trackedEntityType string
var TrackedEntitiesCmd = &cobra.Command{
	Use:   "trackedEntities",
	Short: "Manage tracked entities",
}

func init() {
	TrackedEntitiesCmd.AddCommand(ListCmd, DeleteCmd, ExportCmd)
	TrackedEntitiesCmd.PersistentFlags().StringVar(&program, "program", "", "Program")
	TrackedEntitiesCmd.PersistentFlags().StringVar(&orgUnit, "orgunits", "", "Organization Unit")
	TrackedEntitiesCmd.PersistentFlags().StringVar(&orgUnitMode, "orgunitMode", "SELECTED", "Organization Unit Mode")
	TrackedEntitiesCmd.PersistentFlags().StringVar(&trackedEntityType, "trackedEntityType", "", "Tracked entity type")
	TrackedEntitiesCmd.MarkFlagsMutuallyExclusive("program", "trackedEntityType")
	// TrackedEntitiesCmd.MarkFlagsOneRequired("program", "trackedEntityType")
	_ = TrackedEntitiesCmd.RegisterFlagCompletionFunc("orgunitMode",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := []string{"SELECTED", "CHILDREN", "DESCENDANTS", "ACCESSIBLE", "CAPTURE", "ALL"}
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
}
