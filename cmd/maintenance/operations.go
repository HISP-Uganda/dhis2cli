package maintenance

import (
	"dhis2cli/client"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"reflect"
	"strings"
)

var OperationsCmd = &cobra.Command{
	Use:   "operations",
	Short: "Perform maintenance operations",
	Run: func(c *cobra.Command, args []string) {
		params := make(map[string]interface{})
		if applyAll {
			operationsMap := structToMapStringAny(ops)
			for name, _ := range operationsMap {
				params[name] = "true"
				// Perform operation logic here
			}
		} else {
			params = structToMapStringAny(ops)
		}
		fmt.Printf("Performing maintenance operations...\n%#v\n", params)
		utils.PostResourceAndDisplay(client.Dhis2Client, "maintenance", params, nil, "", "string")
		fmt.Println("Done performing maintenance operations")
	},
}
var applyAll bool

type Operations struct {
	AnalyticsTablesClear                    bool
	AnalyticsTablesAnalyze                  bool
	ExpiredInvitationsClear                 bool
	PeriodPruning                           bool
	ZeroDataValueRemoval                    bool
	SoftDeletedDataValueRemoval             bool
	SoftDeletedProgramStageInstanceRemoval  bool
	SoftDeletedProgramInstanceRemoval       bool
	SoftDeletedTrackedEntityInstanceRemoval bool
	SqlViewsDrop                            bool
	SqlViewsCreate                          bool
	CategoryOptionComboUpdate               bool
	CacheClear                              bool
	OuPathsUpdate                           bool
	AppReload                               bool
}

var ops Operations

func init() {
	OperationsCmd.Flags().BoolVar(&applyAll, "apply-all", false, "Perform all maintenance operations")
	OperationsCmd.Flags().BoolVar(&ops.AnalyticsTablesClear, "analyticsTablesClear", false, "Clear analytics tables")
	OperationsCmd.Flags().BoolVar(&ops.AnalyticsTablesAnalyze, "analyticsTablesAnalyze", false, "Analyze analytics tables")
	OperationsCmd.Flags().BoolVar(&ops.ExpiredInvitationsClear, "expiredInvitationsClear", false, "Remove expired invitations")
	OperationsCmd.Flags().BoolVar(&ops.PeriodPruning, "periodPruning", false, "Prune periods")
	OperationsCmd.Flags().BoolVar(&ops.ZeroDataValueRemoval, "zeroDataValueRemoval", false, "Remove zero data values")
	OperationsCmd.Flags().BoolVar(&ops.SoftDeletedDataValueRemoval, "softDeletedDataValueRemoval", false,
		"Permanently remove soft deleted data values")
	OperationsCmd.Flags().BoolVar(&ops.SoftDeletedProgramStageInstanceRemoval,
		"softDeletedProgramStageInstanceRemoval", false, "Permanently remove soft deleted Program Stage instances")
	OperationsCmd.Flags().BoolVar(&ops.SoftDeletedProgramInstanceRemoval,
		"softDeletedProgramInstanceRemoval", false, "Permanently remove soft deleted Program instances")
	OperationsCmd.Flags().BoolVar(&ops.SoftDeletedTrackedEntityInstanceRemoval, "softDeletedTrackedEntityInstanceRemoval",
		false, "Permanently remove soft deleted tracked entity instances")
	OperationsCmd.Flags().BoolVar(&ops.SqlViewsDrop, "sqlViewsDrop", false, "Drop SQL views")
	OperationsCmd.Flags().BoolVar(&ops.SqlViewsCreate, "sqlViewsCreate", false, "Create SQL views")
	OperationsCmd.Flags().BoolVar(&ops.CategoryOptionComboUpdate, "categoryOptionComboUpdate", false, "Update category option combinations")
	OperationsCmd.Flags().BoolVar(&ops.CacheClear, "cacheClear", false, "Clear application cache")
	OperationsCmd.Flags().BoolVar(&ops.OuPathsUpdate, "ouPathsUpdate", false, "Update Organization Unit Paths")
	OperationsCmd.Flags().BoolVar(&ops.AppReload, "appReload", false, "Reload apps")
}

func structToMapStringAny(s interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	// Ensure the passed argument is a struct
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := typ.Field(i)
			// Check if the field is boolean
			if field.Kind() == reflect.Bool {
				fieldName := toLowerFirstChar(fieldType.Name)
				fieldValue := field.Bool()
				result[fieldName] = fmt.Sprintf("%t", fieldValue)
			}
		}
	}
	return result
}

// toLowerFirstChar converts the first character of a string to lowercase
func toLowerFirstChar(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(string(s[0])) + s[1:]
}
