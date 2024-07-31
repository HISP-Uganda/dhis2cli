package metadata

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var useGist bool
var gistUID string
var gistFieldName string

func init() {
	ListCmd.Flags().StringVarP(&resource, "resource", "r", "", "The metadata resource")
	ListCmd.Flags().BoolVar(&useGist, "gist", false, "Whether to use the Gist API")
	ListCmd.Flags().StringVar(&gistUID, "gist-id", "", "The metadata resource")
	ListCmd.Flags().StringVar(&gistFieldName, "gist-field", "", "The metadata resource")
	_ = ListCmd.MarkFlagRequired("resource")
	_ = ListCmd.RegisterFlagCompletionFunc("resource",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			completions := KnownResources
			return completions, cobra.ShellCompDirectiveNoFileComp
		})
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List metadata",
	Example: `
# list organization units
dhis2 metadata list -r organisationUnits

# list organisationUnits using gist API, -Q options are additional query parameters
# calls /organisationUnits/gist?auto=XS&headless=true&fields=* 
dhis2 metadata list -r organisationUnits --gist -Q auto=XS -f "*" -Q headless=true --format json 

# makes call to /organisationUnits/d893Lz77NrG/children/gist
dhis2 metadata list -r organisationUnits --gist --gist-id d893Lz77NrG --gist-field children --format json 

`,
	Run: func(cmd *cobra.Command, args []string) {
		if resource == "" {
			fmt.Println("Please provide a metadata resource")
			return
		}
		defaultParams := map[string]any{
			"fields":   "id,displayName",
			"order":    "name:asc",
			"paging":   "true",
			"page":     "1",
			"pageSize": "10",
		}
		additionalParams := map[string]any{}
		// excludeKeys := []string{}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		resourcePath := fmt.Sprintf("/%s", resource)
		if useGist {
			if gistUID != "" {
				resourcePath = fmt.Sprintf("%s/%s", resourcePath, gistUID)
			}
			if gistFieldName != "" {
				resourcePath = fmt.Sprintf("%s/%s", resourcePath, gistFieldName)
			}
			resourcePath = fmt.Sprintf("%s/gist", resourcePath)
		}
		utils.FetchResourceAndDisplay2(client.Dhis2Client, resourcePath, params, resource, config.OutputFormat)
	},
}
