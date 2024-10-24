package enrollments

import "github.com/spf13/cobra"

type EnrolmentsParamsConfig struct {
	Program           string `default:"" json:"program"`
	ProgramStatus     string `default:"" json:"programStatus"`
	FollowUp          bool   `default:"true" json:"followUp"`
	TrackedEntity     string `default:"" json:"trackedEntity"`
	TrackedEntityType string `default:"" json:"trackedEntityType"`
	OrgUnits          string `default:"" json:"orgUnits"`
	OrgUnit           string `default:"" json:"orgUnit"`
	OrgUnitMode       string `default:"" json:"orgUnitMode"`
	OuMode            string `default:"" json:"ouMode"`
	UpdatedAfter      string `default:"" json:"updatedAfter"`
	UpdatedBefore     string `default:"" json:"updatedBefore"`
	UpdatedWithin     string `default:"" json:"updatedWithin"`
	EnrolledAfter     string `default:"" json:"enrolledAfter"`
	EnrolledBefore    string `default:"" json:"enrolledBefore"`
	Order             string `default:"" json:"order"`
	Enrollments       string `default:"" json:"enrollments"`
	Enrollment        string `default:"" json:"enrollment"`
	IncludeDeleted    bool   `default:"false" json:"includeDeleted"`
}

var EnrolmentsParams EnrolmentsParamsConfig

var Enrollments = &cobra.Command{
	Use:   "enrollments",
	Short: "Manage enrollments",
}

func init() {
	Enrollments.AddCommand(ListEnrollmentsCmd)
}
