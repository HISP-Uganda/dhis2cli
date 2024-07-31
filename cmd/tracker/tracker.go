package tracker

import (
	"dhis2cli/cmd/tracker/trackedentities"
	"github.com/spf13/cobra"
)

var TrackerCmd = &cobra.Command{
	Use:   "tracker",
	Short: "Manage Tracker Objects",
}

func init() {
	TrackerCmd.AddCommand(trackedentities.TrackedEntities)
}
