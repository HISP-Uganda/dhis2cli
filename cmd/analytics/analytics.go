package analytics

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

type ParamsConfig struct {
	SkipResourceTables   bool `default:"" json:"skipResourceTables"`
	SkipAggregate        bool `default:"" json:"skipAggregate"`
	SkipEvents           bool `default:"" json:"skipEvents"`
	SkipEnrollment       bool `default:"" json:"skipEnrollment"`
	SkipOrgUnitOwnership bool `default:"" json:"skipOrgUnitOwnership"`
	SkipTrackedEntity    bool `default:"true" json:"skipTrackedEntity"`
	SkipOutlier          bool `default:"true" json:"skipOutlier"`
	LastYears            int  `default:"All" json:"lastYears"`
}

var interactive bool

var Params ParamsConfig

var AnalyticsCmd = &cobra.Command{
	Use:   "analytics",
	Short: "Perform analytics tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Perform analytics tasks logic here
		interactive, _ := cmd.Flags().GetBool("interactive")
		if interactive {
			fmt.Println("Running in interactive mode...")
			additionalParams := utils.GetNonDefaultFields(Params)
			params := config.GenerateParams(config.GlobalParams, nil, additionalParams, nil)
			resp, err := client.Dhis2Client.PostResource("resourceTables/analytics", params, nil)
			if err != nil {
				fmt.Printf("Error generating resource tables: %v\n", err)
				return
			}
			// marshal response to HTTPResponse
			var httpResponse HTTPResponse
			err = json.Unmarshal(resp.Body(), &httpResponse)
			if err != nil {
				fmt.Printf("Error unmarshalling analytics generation response: %v\n", err)
				return
			}
			statusEndpoint := strings.TrimPrefix(httpResponse.Response.RelativeNotifierEndpoint, "/api/")
			fmt.Printf("Status endpoint: %s\n", statusEndpoint)
			// monitorJobStatus(statusEndpoint)
			// monitorJobStatusByIndices(statusEndpoint)
			monitorJobStatusByTimeAndMessage(statusEndpoint)
		} else {
			additionalParams := utils.GetNonDefaultFields(Params)
			params := config.GenerateParams(config.GlobalParams, nil, additionalParams, nil)
			utils.PostResourceAndDisplay(client.Dhis2Client, "resourceTables/analytics", params, nil, "", "json")
		}
	},
}

func init() {
	AnalyticsCmd.Flags().BoolVar(&Params.SkipResourceTables, "skipResourceTables", false, "Skip generation of resource tables")
	AnalyticsCmd.Flags().BoolVar(&Params.SkipAggregate, "skipAggregate", false, "Skip generation of aggregate data and completeness data")
	AnalyticsCmd.Flags().BoolVar(&Params.SkipEvents, "skipEvents", false, "Skip generation of event data")
	AnalyticsCmd.Flags().BoolVar(&Params.SkipEnrollment, "skipEnrollment", false, "Skip generation of enrollment data")
	AnalyticsCmd.Flags().BoolVar(&Params.SkipOrgUnitOwnership, "skipOrgUnitOwnership", false, "Skip generation of organization unit ownership data")
	AnalyticsCmd.Flags().IntVar(&Params.LastYears, "lastYears", 0, "Number of last years of data to include")
	AnalyticsCmd.Flags().BoolVar(&Params.SkipTrackedEntity, "skipTrackedEntity", false, "Skip generation of tracked entity data")
	AnalyticsCmd.Flags().BoolVar(&Params.SkipOutlier, "skipOutlier", false, "Skip generation of outlier data")
	AnalyticsCmd.Flags().BoolVar(&interactive, "interactive", false, "Run in interactive mode")

}

func getJobStatus(endPoint string) ([]SubTaskStatus, error) {
	resp, err := client.Dhis2Client.GetResource(endPoint, nil)
	if err != nil {
		return nil, err
	}

	var taskStatus []SubTaskStatus
	err = json.Unmarshal(resp.Body(), &taskStatus)
	if err != nil {
		return nil, err
	}
	return taskStatus, nil
}

func monitorJobStatus(endPoint string) {
	printedUIDs := make(map[string]bool) // Track printed subtasks by UID

	ticker := time.NewTicker(2 * time.Second) // Check every 2 seconds
	defer ticker.Stop()

	for range ticker.C {
		// Fetch the latest job status
		taskStatus, err := getJobStatus(endPoint)
		if err != nil {
			fmt.Printf("Error fetching job status: %v\n", err)
			continue
		}

		// Iterate over each subtask starting from the most recent
		for i := len(taskStatus) - 1; i >= 0; i-- {
			status := taskStatus[i]

			// Print only if the subtask hasn't been printed before
			if !printedUIDs[status.UID] {
				fmt.Printf("Time: %s, Message: %s\n", status.Time, status.Message)
				printedUIDs[status.UID] = true // Mark as printed

				// Stop monitoring if a subtask is completed
				if status.Completed {
					fmt.Println("A task has completed. Stopping monitoring.")
					return
				}
			}
		}
	}
}

func monitorJobStatusByIndices(endPoint string) {
	printedIndices := make(map[int]bool) // Track printed indices for each subtask

	ticker := time.NewTicker(4 * time.Second) // Check every 2 seconds
	defer ticker.Stop()

	for range ticker.C {
		// Fetch the latest job status
		taskStatus, err := getJobStatus(endPoint)
		if err != nil {
			fmt.Printf("Error fetching job status: %v\n", err)
			continue
		}

		// Iterate over each subtask by index
		for i := len(taskStatus) - 1; i >= 0; i-- {
			// Only print the entry if it hasn't been printed before
			if !printedIndices[i] {
				status := taskStatus[i]
				fmt.Printf("Index: %d, Time: %s, Message: %s Completed: %v\n", i, status.Time, status.Message, status.Completed)

				// Mark this index as printed
				printedIndices[i] = true

				// Stop monitoring if a subtask is completed
				if status.Completed {
					fmt.Println("A task has completed. Stopping monitoring.")
					return
				}
			}
		}
	}
}

func monitorJobStatusByTimeAndMessage(endPoint string) {
	printedEntries := make(map[string]bool) // Track printed entries by "Time|Message"

	ticker := time.NewTicker(2 * time.Second) // Check every 2 seconds
	defer ticker.Stop()

	for range ticker.C {
		// Fetch the latest job status
		taskStatus, err := getJobStatus(endPoint)
		if err != nil {
			fmt.Printf("Error fetching job status: %v\n", err)
			continue
		}

		// Iterate over each subtask
		for i := len(taskStatus) - 1; i >= 0; i-- {
			status := taskStatus[i]

			// Create a unique key based on Time and Message
			entryKey := status.Time + "|" + status.Message

			// Print only if this entry hasn't been printed before
			if !printedEntries[entryKey] {
				fmt.Printf("%s,  %s\n", status.Time, status.Message)

				// Mark this entry as printed
				printedEntries[entryKey] = true

				// Stop monitoring if a subtask is completed
				if status.Completed {
					fmt.Println("Analytics Completed!!")
					return
				}
			}
		}
	}
}
