package jobs

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var JobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "View jobs and their reports",
}

var jobId string
var reportMode = "ERRORS"

func init() {
	JobsCmd.AddCommand(ListCmd)
	JobsCmd.AddCommand(ReportCmd)
	ListCmd.Flags().StringVar(&jobId, "uid", "", "The job identifier")
	_ = ListCmd.MarkFlagRequired("uid")
	ReportCmd.Flags().StringVar(&jobId, "uid", "", "The job identifier")
	_ = ReportCmd.MarkFlagRequired("uid")
	ReportCmd.Flags().StringVar(&reportMode, "report-mode", "ERRORS", "The report mode")
	_ = ReportCmd.RegisterFlagCompletionFunc("report-mode",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := []string{"ERRORS", "WARNINGS", "FULL"}
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all jobs",
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := fmt.Sprintf("tracker/jobs/%s", jobId)
		utils.FetchResourceAndDisplay2(client.Dhis2Client, endpoint, nil, "", "json")
	},
}

var ReportCmd = &cobra.Command{
	Use:   "report",
	Short: "View a report for a specific job",
	Run: func(cmd *cobra.Command, args []string) {
		params := map[string]interface{}{
			"reportMode": reportMode,
		}
		endpoint := fmt.Sprintf("tracker/jobs/%s/report", jobId)
		utils.FetchResourceAndDisplay2(client.Dhis2Client, endpoint, params, "", "json")
	},
}
