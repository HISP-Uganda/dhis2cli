package tracker

import (
	"dhis2cli/cmd/tracker/enrollments"
	"dhis2cli/cmd/tracker/events"
	"dhis2cli/cmd/tracker/jobs"
	"dhis2cli/cmd/tracker/trackedentities"
	"dhis2cli/cmd/tracker/trackedentitytypes"
	"dhis2cli/config"
	"github.com/spf13/cobra"
)

var TrackerCmd = &cobra.Command{
	Use:   "tracker",
	Short: "Manage Tracker Objects",
}

func init() {
	TrackerCmd.AddCommand(trackedentities.TrackedEntitiesCmd)
	TrackerCmd.AddCommand(trackedentitytypes.TrackedEntityTypeCmd)
	TrackerCmd.AddCommand(jobs.JobsCmd)
	TrackerCmd.AddCommand(events.Events)
	TrackerCmd.AddCommand(enrollments.Enrollments)
	TrackerCmd.PersistentFlags().StringVarP(&config.GlobalParams.TotalPages,
		"totalPages", "", "false", "Whether to return to return the total number of elements and pages.")
}
