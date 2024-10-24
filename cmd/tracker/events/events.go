package events

import "github.com/spf13/cobra"

type EventParamsConfig struct {
	Program                     string `default:"" json:"program"`
	ProgramStage                string `default:"" json:"programStage"`
	ProgramStatus               string `default:"" json:"programStatus"`
	Filter                      string `default:"" json:"filter"`
	FilterAttributes            string `default:"" json:"filterAttributes"`
	FollowUp                    bool   `default:"true" json:"followUp"`
	TrackedEntity               string `default:"" json:"trackedEntity"`
	OrgUnit                     string `default:"" json:"orgUnit"`
	OrgUnitMode                 string `default:"" json:"orgUnitMode"`
	OuMode                      string `default:"" json:"ouMode"`
	Status                      string `default:"" json:"status"`
	OccurredAfter               string `default:"" json:"occurredAfter"`
	OccurredBefore              string `default:"" json:"occurredBefore"`
	ScheduledAfter              string `default:"" json:"scheduledAfter"`
	ScheduledBefore             string `default:"" json:"scheduledBefore"`
	UpdatedAfter                string `default:"" json:"updatedAfter"`
	UpdatedBefore               string `default:"" json:"updatedBefore"`
	UpdatedWithin               string `default:"" json:"updatedWithin"`
	EnrollmentEnrolledAfter     string `default:"" json:"enrollmentEnrolledAfter"`
	EnrollmentEnrolledBefore    string `default:"" json:"enrollmentEnrolledBefore"`
	EnrollmentOccurredAfter     string `default:"" json:"enrollmentOccurredAfter"`
	EnrollmentOccurredBefore    string `default:"" json:"enrollmentOccurredBefore"`
	DataElementIdScheme         string `default:"" json:"dataElementIdScheme"`
	CategoryOptionComboIdScheme string `default:"" json:"categoryOptionComboIdScheme"`
	OrgUnitIdScheme             string `default:"" json:"orgUnitIdScheme"`
	ProgramIdScheme             string `default:"" json:"programIdScheme"`
	ProgramStageIdScheme        string `default:"" json:"programStageIdScheme"`
	IdScheme                    string `default:"" json:"idScheme"`
	Order                       string `default:"" json:"order"`
	Events                      string `default:"" json:"events"`
	AttributeCategoryCombo      string `default:"" json:"attributeCategoryCombo"`
	AttributeCategoryOptions    string `default:"" json:"attributeCategoryOptions"`
	IncludeDeleted              bool   `default:"false" json:"includeDeleted"`
	AssignedUserMode            string `default:"" json:"assignedUserMode"`
	AssignedUsers               string `default:"" json:"assignedUsers"`
}

var EventParams EventParamsConfig

var Events = &cobra.Command{
	Use:   "events",
	Short: "Manage events",
}
var OuMode = []string{"SELECTED", "CHILDREN", "DESCENDANTS", "ACCESSIBLE", "CAPTURE", "ALL"}
var EventStatusOptions = []string{"ACTIVE", "COMPLETED", "VISITED", "SCHEDULE", "OVERDUE", "SKIPPED"}
var ProgramStatusOptions = []string{"ACTIVE", "COMPLETED", "CANCELED"}

func init() {
	Events.AddCommand(ListEventsCmd)
	//Events.AddCommand(CreateCmd)
	//Events.AddCommand(DeleteCmd)
	//Events.AddCommand(UpdateCmd)
	//Events.AddCommand(QueryCmd)
}
