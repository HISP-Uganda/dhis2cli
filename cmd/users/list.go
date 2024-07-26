package users

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		params := map[string]string{
			"paging":          config.Paging,
			"fields":          config.Fields,
			"order":           "firstName:asc,surname:asc",
			"userOrgUnits":    "true",
			"includeChildren": "true",
		}
		if config.Query != "" {
			params["query"] = fmt.Sprintf("%s", config.Query)
		}
		if config.Filter != "" {
			params["filter"] = config.Filter
		}
		if config.Paging == "true" {
			params["page"] = fmt.Sprintf("%d", config.Page)
			params["pageSize"] = fmt.Sprintf("%d", config.PageSize)
		}
		// XXX Search by OU: organisationUnits.id:in:[akV6429SUqu]
		// XXX Search by name
		resp, err := client.Dhis2Client.GetResource(
			"/users", params)
		if err != nil {
			fmt.Printf("Error fetching users: %v\n", err)
			return
		}
		// fmt.Println(string(resp.Body()))
		// var usersList []models.User
		var users map[string]any
		// var usersList []map[string]string
		err = json.Unmarshal(resp.Body(), &users)
		if err != nil {
			fmt.Printf("Error unmarshalling users list: %v\n", err)
			return
		}

		//usersList, err := utils.ConvertToMapStringStringSlice(users["users"])
		//if err != nil {
		//	fmt.Printf("Error converting users to slice of maps: %v\n", err)
		//	return
		//}
		//
		//fmt.Printf("%s\n", utils.CreateDataFrameFromMap(usersList))
		// utils.DisplayTable(users["users"])
		// Display the table
		err = utils.DisplayTable(users["users"])
		// err = utils.DisplayTable2(users["users"], true)
		if err != nil {
			fmt.Println("Error:", err)
		}
		prettyJSON, err := json.MarshalIndent(users, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("\n%s\n", string(prettyJSON))

	},
}

//func init() {
//	ListCmd.Flags().StringVarP(&paging, "paging", "p", "true", "Whether to set pagination")
//	ListCmd.Flags().StringVarP(&fields, "fields", "f", "id,displayName", "Fields to display")
//	ListCmd.Flags().StringVarP(&filter, "filter", "F", "", "Filters to apply")
//	ListCmd.Flags().StringVarP(&query, "query", "q", "", "Query term")
//}
