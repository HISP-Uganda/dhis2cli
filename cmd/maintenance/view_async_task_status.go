package maintenance

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var taskCategoryIDs = []string{
	"ANALYTICS_TABLE",
	"RESOURCE_TABLE",
	"MONITORING",
	"DATAVALUE_IMPORT",
	"EVENT_IMPORT",
	"ENROLLMENT_IMPORT",
	"TEI_IMPORT",
	"METADATA_IMPORT",
	"DATA_INTEGRITY",
}

var taskCategory string
var taskId string

// ViewAsyncTasksStatusCmd is the command for viewing the status of asynchronous tasks
var ViewAsyncTasksStatusCmd = &cobra.Command{
	Use:   "viewAsyncTaskStatus",
	Short: "View the status of a asynchronous tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// View async task status logic here
		resourcePath := fmt.Sprintf("system/tasks/%s", taskCategory)
		if taskId != "" {
			resourcePath = fmt.Sprintf("%s/%s", resourcePath, taskId)
		}
		utils.FetchResourceAndDisplay2(client.Dhis2Client, resourcePath, nil, "", "json")
	},
}

func init() {
	ViewAsyncTasksStatusCmd.Flags().StringVar(&taskCategory, "task-category", "", "Task Category Identifier")
	ViewAsyncTasksStatusCmd.Flags().StringVar(&taskId, "task-id", "", "Task Identifier")
	_ = ViewAsyncTasksStatusCmd.MarkFlagRequired("task-category")
	_ = ViewAsyncTasksStatusCmd.RegisterFlagCompletionFunc("task-category",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := taskCategoryIDs
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
}
