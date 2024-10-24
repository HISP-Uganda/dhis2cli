package enrollments

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var ListEnrollmentsCmd = &cobra.Command{
	Use:   "list",
	Short: "List enrollments",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement fetching and displaying enrollments
		defaultParams := map[string]any{
			"fields":     "*",
			"order":      "createdAt:desc",
			"paging":     "true",
			"page":       "1",
			"pageSize":   "10",
			"totalPages": config.GlobalParams.TotalPages,
		}
		additionalParams := utils.GetNonDefaultFields(EnrolmentsParams)
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		resource := ""
		switch config.OutputFormat {
		case "csv", "table":
			resource = "instances"
		}

		utils.FetchResourceAndDisplay2(client.Dhis2Client, "tracker/enrollments", params, resource, config.OutputFormat)
	},
}

func init() {
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.Program, "program", "", "Filter enrollments by program")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.OrgUnits, "orgUnits", "", "Comma-separated list of organisation unit UIDs")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.OrgUnit, "orgUnit", "", "Semicolon-separated list of organisation units UIDs")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.OrgUnitMode, "orgUnitMode", "", "Mode for orgUnit")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.OuMode, "ouMode", "", "Mode for orgUnit")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.ProgramStatus, "programStatus", "", "Filter enrollments by program status")
	ListEnrollmentsCmd.Flags().BoolVar(&EnrolmentsParams.FollowUp, "followUp", true, "Include follow-up enrollments")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.UpdatedAfter,
		"updatedAfter", "", "Only enrollments updated after this date")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.UpdatedWithin,
		"updatedWithin", "", "Only enrollments updated since given duration")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.EnrolledBefore,
		"enrolledBefore", "", "Only enrollments older than this date")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.EnrolledAfter,
		"enrolledAfter", "", "Only enrollments newer than this date")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.TrackedEntity, "trackedEntity", "", "Identifier of tracked entity")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.TrackedEntityType, "trackedEntityType", "", "Identifier of tracked entity type")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.Order, "order", "", "Comma-separated list of property name, attribute or data element UID and sort direction pairs in format propName:sortDirection")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.Enrollments, "enrollments", "", "Comma-separated list of enrollments UIDs")
	ListEnrollmentsCmd.Flags().StringVar(&EnrolmentsParams.Enrollment, "enrollment", "", "Semicolon-separated list of enrollment UIDs")
	ListEnrollmentsCmd.Flags().BoolVar(&EnrolmentsParams.IncludeDeleted,
		"includeDeleted", false, "When true, soft deleted enrollments will be included in your query result.")

}
