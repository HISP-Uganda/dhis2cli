package data

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

type TemplateParamsConfig struct {
	Dataset             string `default:"" json:"dataset"`
	Period              string `default:"" json:"period"`
	OrgUnit             string `default:"" json:"orgUnit"`
	Comment             string `default:"Yes" json:"comment"`
	OrgUnitIdScheme     string `default:"" json:"orgUnitIdScheme"`
	DataElementIdScheme string `default:"" json:"dataElementIdScheme"`
}

var templateParams TemplateParamsConfig

var DataValueSetTemplateCmd = &cobra.Command{
	Use:   "dataValueSetTemplate",
	Short: "Generate data value set template",
	Run: func(cmd *cobra.Command, args []string) {
		additionalParams := utils.GetNonDefaultFields(templateParams)
		params := config.GenerateParams(config.GlobalParams, nil, additionalParams, nil)
		endpoint := fmt.Sprintf("dataSets/%s/dataValueSet", templateParams.Dataset)
		utils.FetchResourceAndDisplay2(client.Dhis2Client, endpoint, params, "", "json")
	},
}

func init() {
	DataValueSetTemplateCmd.Flags().StringVar(&templateParams.Dataset, "dataset", "", "Dataset UID")
	DataValueSetTemplateCmd.Flags().StringVar(&templateParams.Period, "period", "", "Period to use, will be included without any checks.")
	DataValueSetTemplateCmd.Flags().StringVar(&templateParams.OrgUnit, "orgUnit", "",
		"Organisation unit to use, supports multiple orgUnits, both id and code can be used.")
	DataValueSetTemplateCmd.Flags().StringVar(&templateParams.Comment, "comment", "", "Comment for the template.")
	DataValueSetTemplateCmd.Flags().StringVar(&templateParams.OrgUnitIdScheme,
		"orgUnitIdScheme", "", "Organisation unit scheme to use")
	DataValueSetTemplateCmd.Flags().StringVar(&templateParams.DataElementIdScheme,
		"dataElementIdScheme", "", "Data element scheme to use")

	_ = DataValueSetTemplateCmd.MarkFlagRequired("dataset")
}
