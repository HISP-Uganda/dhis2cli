package maintenance

import "github.com/spf13/cobra"

var MaintenanceCmd = &cobra.Command{
	Use:   "maintenance",
	Short: "Perform maintenance tasks",
	//Run: func(cmd *cobra.Command, args []string) {
	//    // Perform maintenance tasks logic here
	//},
}

func init() {
	MaintenanceCmd.AddCommand(
		ViewAsyncTasksStatusCmd,
		ViewAsyncTaskSummariesCmd,
		OperationsCmd,
		ResourceTablesCmd)
}
