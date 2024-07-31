package maintenance

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var ViewAsyncTaskSummariesCmd = &cobra.Command{
	Use:   "viewAsyncTaskSummaries",
	Short: "View the summary of asynchronous tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// View async task summary logic here
		resourcePath := fmt.Sprintf("system/taskSummaries/%s/%s", taskCategory, taskId)

		utils.FetchResourceAndDisplay2(client.Dhis2Client, resourcePath, nil, "", "json")
	},
}

func init() {
	ViewAsyncTaskSummariesCmd.Flags().StringVar(&taskCategory, "task-category", "", "Task Category Identifier")
	ViewAsyncTaskSummariesCmd.Flags().StringVar(&taskId, "task-id", "", "Task Identifier")
	_ = ViewAsyncTaskSummariesCmd.RegisterFlagCompletionFunc("task-category",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := taskCategoryIDs
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
	_ = ViewAsyncTaskSummariesCmd.MarkFlagRequired("task-category")
	_ = ViewAsyncTaskSummariesCmd.MarkFlagRequired("task-id")
}
