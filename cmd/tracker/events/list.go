package events

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"github.com/spf13/cobra"
)

var ListEventsCmd = &cobra.Command{
	Use:   "list",
	Short: "List events",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement fetching and displaying events
		// log.Infof("The parameters for query are %v, Full Conf: %v", utils.GetNonDefaultFields(EventParams), EventParams)
		defaultParams := map[string]any{
			"fields":     "*",
			"order":      "createdAt:desc",
			"paging":     "true",
			"page":       "1",
			"pageSize":   "10",
			"totalPages": config.GlobalParams.TotalPages,
		}
		additionalParams := utils.GetNonDefaultFields(EventParams)
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		resource := ""
		switch config.OutputFormat {
		case "csv", "table":
			resource = "instances"
		}

		utils.FetchResourceAndDisplay2(client.Dhis2Client, "tracker/events", params, resource, config.OutputFormat)
	},
}

func init() {
	ListEventsCmd.Flags().StringVar(&EventParams.Program, "program", "", "Identifier of program")
	ListEventsCmd.Flags().StringVar(&EventParams.OrgUnit, "orgUnit", "", "Identifier of organisation unit")
	ListEventsCmd.Flags().StringVar(&EventParams.TrackedEntity, "trackedEntity", "", "Identifier of tracked entity")
	ListEventsCmd.Flags().StringVar(&EventParams.ProgramStage, "programStage", "", "Identifier of program stage")
	ListEventsCmd.Flags().StringVar(&EventParams.Status, "status", "", "Status of event")
	ListEventsCmd.Flags().StringVar(&EventParams.Filter, "filter", "", "Comma separated values of data element filters")
	ListEventsCmd.Flags().StringVar(&EventParams.FilterAttributes, "filterAttributes", "", "Comma separated values of attribute filters")
	ListEventsCmd.Flags().BoolVar(&EventParams.FollowUp, "followUp", true, "Include follow-up events")
	ListEventsCmd.Flags().StringVar(&EventParams.OrgUnitMode, "orgUnitMode", "", "Mode for orgUnit")
	ListEventsCmd.Flags().StringVar(&EventParams.OuMode, "ouMode", "", "Mode for orgUnit")
	ListEventsCmd.Flags().StringVar(&EventParams.ScheduledAfter, "scheduledAfter", "", "Events scheduled after this date")
	ListEventsCmd.Flags().StringVar(&EventParams.ScheduledBefore, "scheduledBefore", "", "Events scheduled before this date")
	ListEventsCmd.Flags().StringVar(&EventParams.UpdatedAfter, "updatedAfter", "", "Events updated after this date")
	ListEventsCmd.Flags().StringVar(&EventParams.UpdatedBefore, "updatedBefore", "", "Events updated before this date")
	ListEventsCmd.Flags().StringVar(&EventParams.UpdatedWithin, "updatedWithin", "", "Events updated within this date range")
	ListEventsCmd.Flags().StringVar(&EventParams.EnrollmentEnrolledAfter,
		"enrollmentEnrolledAfter", "", "Start date and time for enrollment in the given program")
	ListEventsCmd.Flags().StringVar(&EventParams.EnrollmentEnrolledBefore,
		"enrollmentEnrolledBefore", "", "End date and time for enrollment in the given program")
	ListEventsCmd.Flags().StringVar(&EventParams.EnrollmentOccurredAfter,
		"enrollmentOccurredAfter", "", "Start date and time for occurred in the given program stage")
	ListEventsCmd.Flags().StringVar(&EventParams.EnrollmentOccurredBefore,
		"enrollmentOccurredBefore", "", "End date and time for occurred in the given program stage")
	ListEventsCmd.Flags().StringVar(&EventParams.DataElementIdScheme,
		"dataElementIdScheme", "", "Data element identifier scheme")
	ListEventsCmd.Flags().StringVar(&EventParams.CategoryOptionComboIdScheme,
		"categoryOptionComboIdScheme", "", "Category Option Combo ID scheme to use for export")
	ListEventsCmd.Flags().StringVar(&EventParams.OrgUnitIdScheme, "orgUnitIdScheme", "", "Org Unit ID scheme to use for export")
	ListEventsCmd.Flags().StringVar(&EventParams.ProgramIdScheme,
		"trackedEntityIdScheme", "", "Program ID scheme to use for export")
	ListEventsCmd.Flags().StringVar(&EventParams.ProgramStageIdScheme,
		"programStageIdScheme", "", "Program Stage ID scheme to use for export")
	ListEventsCmd.Flags().StringVar(&EventParams.IdScheme,
		"idScheme", "", "Set data element, category option combo, orgUnit, program and program stage at once.")
	ListEventsCmd.Flags().StringVar(&EventParams.Order, "order", "", "Comma-separated list of property name, attribute or data element UID and sort direction pairs in format propName:sortDirection")
	ListEventsCmd.Flags().StringVar(&EventParams.Events, "events", "", "Comma-separated list of event UIDs")
	ListEventsCmd.Flags().StringVar(&EventParams.AttributeCategoryCombo, "attributeCategoryCombo", "",
		"Attribute category combo identifier")
	ListEventsCmd.Flags().StringVar(&EventParams.AttributeCategoryOptions, "attributeCategoryOptions", "",
		"Comma-separated attribute category option identifiers.")
	ListEventsCmd.Flags().BoolVar(&EventParams.IncludeDeleted,
		"includeDeleted", false, "When true, soft deleted events will be included in your query result.")
	ListEventsCmd.Flags().StringVar(&EventParams.AssignedUserMode, "assignedUserMode", "", "Assigned user selection mode")
	ListEventsCmd.Flags().StringVar(&EventParams.AssignedUsers, "assignedUsers", "",
		"Comma-separated list of user UIDs to filter based on events assigned to the users.")

	_ = ListEventsCmd.RegisterFlagCompletionFunc("orgUnitMode",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := OuMode
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ListEventsCmd.RegisterFlagCompletionFunc("status",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := EventStatusOptions
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ListEventsCmd.RegisterFlagCompletionFunc("programStatus",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := ProgramStatusOptions
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ListEventsCmd.RegisterFlagCompletionFunc("idScheme",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := []string{"CODE", "UID"}
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ListEventsCmd.RegisterFlagCompletionFunc("idScheme",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := []string{"CODE", "UID"}
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ListEventsCmd.RegisterFlagCompletionFunc("programIdScheme",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := []string{"CODE", "UID"}
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ListEventsCmd.RegisterFlagCompletionFunc("dataElementIdScheme",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := []string{"CODE", "UID"}
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ListEventsCmd.RegisterFlagCompletionFunc("orgUnitIdScheme",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := []string{"CODE", "UID"}
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ListEventsCmd.RegisterFlagCompletionFunc("categoryOptionComboIdScheme",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := []string{"CODE", "UID"}
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ListEventsCmd.RegisterFlagCompletionFunc("programStageIdScheme",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := []string{"CODE", "UID"}
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ListEventsCmd.RegisterFlagCompletionFunc("assignedUserMode",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := []string{"CURRENT", "PROVIDED", "NONE", "ANY"}
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	ListEventsCmd.MarkFlagsMutuallyExclusive("updatedAfter", "updatedWithin")
	ListEventsCmd.MarkFlagsMutuallyExclusive("updatedBefore", "updatedWithin")
}
