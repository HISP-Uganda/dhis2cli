package tracker

import (
	"dhis2cli/cmd/tracker/jobs"
	"dhis2cli/cmd/tracker/trackedentities"
	"dhis2cli/cmd/tracker/trackedentitytypes"
	"dhis2cli/config"
	"github.com/spf13/cobra"
)

var TotalPages string
var TrackerCmd = &cobra.Command{
	Use:   "tracker",
	Short: "Manage Tracker Objects",
}

func init() {
	TrackerCmd.AddCommand(trackedentities.TrackedEntitiesCmd)
	TrackerCmd.AddCommand(trackedentitytypes.TrackedEntityTypeCmd)
	TrackerCmd.AddCommand(jobs.JobsCmd)
	TrackerCmd.PersistentFlags().StringVarP(&config.GlobalParams.TotalPages,
		"totalPages", "", "false", "Whether to return to return the total number of elements and pages.")
}
