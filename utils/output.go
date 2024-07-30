package utils

import (
	"dhis2cli/client"
	"encoding/json"
	"fmt"
)

func PrintResponse(responseMap map[string]any, pretty bool) (string, error) {
	if pretty {
		prettyJSON, err := json.MarshalIndent(responseMap, "", "  ")
		if err != nil {
			return "", err
		}
		return string(prettyJSON), nil
	} else {
		retJson, err := json.Marshal(responseMap)
		if err != nil {
			return "", err
		}
		return string(retJson), nil
	}
}

// FetchResourceAndDisplay is a utility function to fetch a resource and display it
func FetchResourceAndDisplay(client *client.Client, endpoint string, params map[string]string, resource, outputFormat string) {
	resp, err := client.GetResource(endpoint, params)
	if err != nil {
		fmt.Printf("Error fetching resource from %s: %v\n", endpoint, err)
		return
	}

	var responseMap map[string]any
	err = json.Unmarshal(resp.Body(), &responseMap)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return
	}

	switch outputFormat {
	case "table":
		err = DisplayTable(responseMap[resource])
		if err != nil {
			fmt.Println("Error displaying table:", err)
		}
	case "json":
		prettyJson, err := PrintResponse(responseMap, true)
		if err != nil {
			fmt.Println("Error pretty printing JSON:", err)
			return
		}
		fmt.Println(prettyJson)
	default:
		fmt.Println("Unsupported output format:", outputFormat)
	}
}
