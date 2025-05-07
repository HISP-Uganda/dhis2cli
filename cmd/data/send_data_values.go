package data

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var payload string
var payloadFile string

type ImportParameters struct {
	DataElementIdScheme          string `default:"" json:"dataElementIdScheme"`
	OrgUnitIdScheme              string `default:"" json:"orgUnitIdScheme"`
	AttributeOptionComboIdScheme string `default:"" json:"attributeOptionComboIdScheme"`
	CategoryOptionComboIdScheme  string `default:"" json:"categoryOptionComboIdScheme"`
	DataSetIdScheme              string `default:"" json:"dataSetIdScheme"`
	CategoryIdScheme             string `default:"" json:"CategoryIdScheme"`
	CategoryOptionIdScheme       string `default:"" json:"categoryOptionIdScheme"`
	IdScheme                     string `default:"" json:"idScheme"`
	Preheat                      bool   `default:"" json:"preheat"`
	DryRun                       bool   `default:"" json:"dryRun"`
	ImportStrategy               string `default:"" json:"importStrategy"`
	SkipExistingCheck            bool   `default:"" json:"skipExistingCheck"`
	SkipAudit                    bool   `default:"" json:"skipAudit"`
	Async                        bool   `default:"" json:"async"`
	Force                        bool   `default:"" json:"force"`
	DataSet                      string `default:"" json:"dataSet"`
}

var ImportParams ImportParameters

var UidSchemes = []string{"id", "code", "name"}

var SendDataValuesCmd = &cobra.Command{
	Use:   "sendDataValues",
	Short: "Send data values",
	Run: func(cmd *cobra.Command, args []string) {
		// Send data values logic here
		var toValidate string
		var err error
		switch {
		case payloadFile != "":
			// Read contents from the specified payload file
			toValidate, err = utils.ReadFileToString(payloadFile)
			if err != nil {
				fmt.Printf("Error reading payload file: %v\n", err)
				return
			}
		case payload != "":
			// Use the provided payload string
			toValidate = payload
		default:
			fmt.Println("No payload provided.")
			return
		}

		additionalParams := utils.GetNonDefaultFields(ImportParams)
		params := config.GenerateParams(config.GlobalParams, nil, additionalParams, nil)

		// fmt.Printf("%v\n", toValidate)
		utils.PostResourceAndDisplay(client.Dhis2Client, "dataValueSets", params, toValidate, "", "json")
	},
}

func init() {
	SendDataValuesCmd.Flags().StringVar(&payload, "payload", "", "The payload to send")
	SendDataValuesCmd.Flags().StringVar(&payloadFile, "payload-file", "", "The file with payload to send")
	SendDataValuesCmd.MarkFlagsMutuallyExclusive("payload", "payload-file")
	_ = SendDataValuesCmd.MarkFlagFilename("payload-file", ".json", ".csv")

	SendDataValuesCmd.Flags().StringVar(&ImportParams.DataElementIdScheme, "dataElementIdScheme", "",
		fmt.Sprintf("Property of the data element object to use to map the data values (one of: %s)", strings.Join(UidSchemes, ", ")))

	SendDataValuesCmd.Flags().StringVar(&ImportParams.OrgUnitIdScheme, "orgUnitIdScheme", "",
		fmt.Sprintf("Property of the org unit object to use to map the data values (one of: %s)", strings.Join(UidSchemes, ", ")))

	SendDataValuesCmd.Flags().StringVar(&ImportParams.AttributeOptionComboIdScheme, "attributeOptionComboIdScheme", "",
		fmt.Sprintf("Property of the attribute option combo object to use to map the data values (one of: %s)", strings.Join(UidSchemes, ", ")))

	SendDataValuesCmd.Flags().StringVar(&ImportParams.CategoryOptionComboIdScheme, "categoryOptionComboIdScheme", "",
		fmt.Sprintf("Property of the category option combo object to use to map the data values. (one of: %s)", strings.Join(UidSchemes, ", ")))

	SendDataValuesCmd.Flags().StringVar(&ImportParams.DataSetIdScheme, "dataSetIdScheme", "",
		fmt.Sprintf("Property of the data set object to use to map the data values (one of: %s)", strings.Join(UidSchemes, ", ")))

	SendDataValuesCmd.Flags().StringVar(&ImportParams.CategoryIdScheme, "categoryIdScheme", "",
		fmt.Sprintf("Property of the category object to use to map the data values (ADX only) (one of: %s)", strings.Join(UidSchemes, ", ")))

	SendDataValuesCmd.Flags().StringVar(&ImportParams.CategoryOptionIdScheme, "categoryOptionIdScheme", "",
		fmt.Sprintf("Property of the category option object to use to map the data values (ADX only) (one of: %s)", strings.Join(UidSchemes, ", ")))

	SendDataValuesCmd.Flags().StringVar(&ImportParams.IdScheme, "idScheme", "",
		fmt.Sprintf("Property of any of the above objects if they are not specified, to use to map the data values. (one of: %s)", strings.Join(UidSchemes, ", ")))

	SendDataValuesCmd.Flags().BoolVar(&ImportParams.Preheat, "preheat", false, "Indicates whether to preload metadata caches before starting to import data values.")
	SendDataValuesCmd.Flags().BoolVar(&ImportParams.DryRun, "dryRun", false, "Whether to save changes on the server or just return the import summary.")
	SendDataValuesCmd.Flags().StringVar(&ImportParams.ImportStrategy,
		"importStrategy", "", "Save objects of all, new or update import status on the server.")
	SendDataValuesCmd.Flags().BoolVar(&ImportParams.SkipExistingCheck,
		"skipExistingCheck", false, "Skip checks for existing data values. Improves performance. .")
	SendDataValuesCmd.Flags().BoolVar(&ImportParams.SkipAudit,
		"skipAudit", false, "Skip audit, meaning audit values will not be generated. Improves performance at the cost of ability to audit changes.")
	SendDataValuesCmd.Flags().BoolVar(&ImportParams.Async,
		"async", false, "Asynchronously send data values.")
	SendDataValuesCmd.Flags().BoolVar(&ImportParams.Force,
		"force", false, "Force import, even if the import is not possible due to missing dependencies.")
	SendDataValuesCmd.Flags().StringVar(&ImportParams.DataSet,
		"dataSet", "", "Provide the data set ID for CSV import where the ID cannot be provided in the file itself")

	_ = SendDataValuesCmd.RegisterFlagCompletionFunc("importStrategy",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := utils.ImportStrategy
			return completions, cobra.ShellCompDirectiveNoFileComp
		})

}
