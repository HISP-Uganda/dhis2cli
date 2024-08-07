package metadata

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var payload string
var payloadFile string
var ResourcesSingular = []string{
	"aggregateDataExchange",
	"analyticsTableHook",
	"apiToken",
	"attribute",
	"category",
	"categoryCombo",
	"categoryOption",
	"categoryOptionCombo",
	"categoryOptionGroup",
	"categoryOptionGroupSet",
	"constant",
	"dashboard",
	"dashboardItem",
	"dataApprovalLevel",
	"dataApprovalWorkflow",
	"dataElement",
	"dataElementGroup",
	"dataElementGroupSet",
	"dataElementOperand",
	"dataEntryForm",
	"dataSet",
	"dataSetNotificationTemplate",
	"dataStore",
	"document",
	"eventChart",
	"eventFilter",
	"eventReport",
	"eventVisualization",
	"externalFileResource",
	"externalMapLayer",
	"fileResource",
	"icon",
	"indicator",
	"indicatorGroup",
	"indicatorGroupSet",
	"indicatorType",
	"interpretation",
	"jobConfiguration",
	"legendSet",
	"map",
	"mapView",
	"messageConversation",
	"metadataVersion",
	"minMaxDataElement",
	"oAuth2Client",
	"option",
	"optionGroup",
	"optionGroupSet",
	"optionSet",
	"organisationUnit",
	"organisationUnitGroup",
	"organisationUnitGroupSet",
	"organisationUnitLevel",
	"predictor",
	"predictorGroup",
	"program",
	"programDataElement",
	"programIndicator",
	"programIndicatorGroup",
	"programNotificationTemplate",
	"programRule",
	"programRuleAction",
	"programRuleVariable",
	"programSection",
	"programStage",
	"programStageSection",
	"programTrackedEntityAttributeGroup",
	"proposal",
	"pushAnalysis",
	"relationship",
	"relationshipType",
	"report",
	"section",
	"smsCommand",
	"sqlView",
	"trackedEntityAttribute",
	"trackedEntityInstanceFilter",
	"trackedEntityType",
	"user",
	"userGroup",
	"userRole",
	"validationNotificationTemplate",
	"validationResult",
	"validationRule",
	"validationRuleGroup",
	"visualization",
}

// ValidateCmd is a command to validate payloads
var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate metadata payloads",
	Run: func(cmd *cobra.Command, args []string) {
		// Validate metadata objects logic here
		// use the /api/schemas/resource endpoint
		resourcePath := fmt.Sprintf("schemas/%s", resource)
		var toValidate string

		if payloadFile != "" {
			// Read contents from the specified payload file
			var err error
			toValidate, err = utils.ReadFileToString(payloadFile)
			if err != nil {
				fmt.Printf("Error reading payload file: %v\n", err)
				return
			}
		} else if payload != "" {
			// Use the provided payload string
			toValidate = payload
		} else {
			fmt.Println("No payload provided.")
			return
		}
		utils.PostResourceAndDisplay(client.Dhis2Client, resourcePath, nil, toValidate, "", "json")
	},
}

func init() {
	ValidateCmd.Flags().StringVar(&payload, "payload", "", "The payload to validate")
	ValidateCmd.Flags().StringVar(&payloadFile, "payload-file", "", "The file with payload to validate")
	ValidateCmd.Flags().StringVarP(&resource, "resource", "r", "", "The metadata resource")
	ValidateCmd.MarkFlagsMutuallyExclusive("payload", "payload-file")
	_ = ValidateCmd.MarkFlagFilename("payload-file", ".json")
	_ = ValidateCmd.MarkFlagRequired("resource")
	_ = ValidateCmd.RegisterFlagCompletionFunc("resource",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := ResourcesSingular
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
}
