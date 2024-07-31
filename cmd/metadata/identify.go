package metadata

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var UID string
var IdentityCmd = &cobra.Command{
	Use:   "identity",
	Short: "Identify metadata object by id",
	Run: func(cmd *cobra.Command, args []string) {
		if UID == "" || !utils.IsValidDHIS2UID(UID) {
			fmt.Printf("Please provide a valid DHIS2 UID: Provided '%s'\n", UID)
			return
		}
		defaultParams := map[string]any{
			"fields": "id,name,code,href",
		}
		additionalParams := map[string]any{}
		// excludeKeys := []string{}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, nil)
		utils.FetchResourceAndDisplay2(client.Dhis2Client, fmt.Sprintf("/identifiableObjects/%s", UID), params, "", "json")
	},
}

func init() {
	IdentityCmd.PersistentFlags().StringVarP(&UID, "id", "", "", "UID")
	_ = IdentityCmd.MarkFlagRequired("id")
}
