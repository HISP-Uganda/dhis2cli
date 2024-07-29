package users

import (
	"dhis2cli/client"
	"dhis2cli/config"
	"dhis2cli/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

var format = "table"

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		defaultParams := map[string]string{
			"fields":          "id,displayName,access,email,username,lastLogin",
			"order":           "firstName:asc,surname:asc",
			"userOrgUnits":    "true",
			"includeChildren": "true",
			"selfRegistered":  "false",
		}
		additionalParams := map[string]string{}
		excludeKeys := []string{}
		params := config.GenerateParams(config.GlobalParams, defaultParams, additionalParams, excludeKeys)

		resp, err := client.Dhis2Client.GetResource(
			"/users", params)
		if err != nil {
			fmt.Printf("Error fetching users: %v\n", err)
			return
		}
		// fmt.Println(string(resp.Body()))
		// var usersList []models.User
		var responseMap map[string]any
		// var usersList []map[string]string
		err = json.Unmarshal(resp.Body(), &responseMap)
		if err != nil {
			fmt.Printf("Error unmarshalling users list: %v\n", err)
			return
		}

		//fmt.Printf("%s\n", utils.CreateDataFrameFromMap(usersList))
		// utils.DisplayTable(responseMap["users"])
		// Display the table
		switch format {
		case "table":
			err = utils.DisplayTable(responseMap["users"])
			// err = utils.DisplayTable2(users["users"], true)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "json":
			prettyJson, err := utils.PrintResponse(responseMap, true)
			if err != nil {
				fmt.Println("Error pretty printing JSON:", err)
				return
			}
			fmt.Println(prettyJson)
		}

	},
}

func init() {
	ListCmd.Flags().StringVarP(&format, "format", "", "table", "Response format: table/json/csv")
}
