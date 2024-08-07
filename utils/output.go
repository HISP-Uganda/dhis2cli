package utils

import (
	"bytes"
	"dhis2cli/client"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

func PrintResponse(responseMap any, pretty bool) (string, error) {
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
func FetchResourceAndDisplay(client *client.Client, endpoint string, params map[string]any, resource, outputFormat string) {
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
		if resource != "" {
			err = DisplayTable(responseMap[resource])

		} else {
			// XXX: for direct arrays of objects, resource should be empty
			err = DisplayTable(responseMap)
		}
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

// FetchResourceAndDisplay2 is a utility function to fetch a resource and display it
func FetchResourceAndDisplay2(client *client.Client, endpoint string, params map[string]any, resource, outputFormat string) {
	resp, err := client.GetResource(endpoint, params)
	if err != nil {
		fmt.Printf("Error fetching resource from %s: %v\n", endpoint, err)
		return
	}
	if outputFormat == "string" {
		fmt.Println(string(resp.Body()))
		return
	}

	var responseMap interface{}
	err = json.Unmarshal(resp.Body(), &responseMap)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return
	}

	switch outputFormat {
	case "table":
		var dataToDisplay interface{}
		if resource != "" {
			if resMap, ok := responseMap.(map[string]interface{}); ok {
				dataToDisplay = resMap[resource]
			} else {
				fmt.Println("Resource key specified, but response is not a map")
				return
			}
		} else {
			dataToDisplay = responseMap
		}
		if fields, ok := params["fields"]; ok {
			fieldsStr := strings.Split(fields.(string), ",")
			if err := DisplayOrderedTable(dataToDisplay, fieldsStr); err != nil {
				fmt.Println("Error displaying table:", err)
			}

		} else {
			if err := DisplayTable(dataToDisplay); err != nil {
				fmt.Println("Error displaying table:", err)
			}
		}

	case "json":
		prettyJson, err := PrintResponse(responseMap, true)
		if err != nil {
			fmt.Println("Error pretty printing JSON:", err)
			return
		}
		fmt.Println(prettyJson)
	case "csv":
		var dataToDisplay interface{}
		if resource != "" {
			if resMap, ok := responseMap.(map[string]interface{}); ok {
				dataToDisplay = resMap[resource]
			} else {
				fmt.Println("Resource key specified, but response is not a map")
				return
			}
		} else {
			dataToDisplay = responseMap
		}
		// csvString, err := AnyToCSV(dataToDisplay)
		// check for fields key in params argument to this function if not present use config.GlobalParams.Fields
		if fields, ok := params["fields"]; ok {
			fieldsStr := strings.Split(fields.(string), ",")
			csvString, err := AnyToCSVWithOrder(dataToDisplay, fieldsStr)
			if err != nil {
				fmt.Printf("Error printing CSV %v\n", err)
				return
			}
			fmt.Println(csvString)
		} else {
			csvString, err := AnyToCSV(dataToDisplay)
			if err != nil {
				fmt.Printf("Error printing CSV %v\n", err)
				return
			}
			fmt.Println(csvString)
		}

	default:
		fmt.Println("Unsupported output format:", outputFormat)
	}
}

// PostResourceAndDisplay posts data to a resource and displays the response
func PostResourceAndDisplay(client *client.Client, endpoint string, params map[string]any, data any, resource, outputFormat string) {
	resp, err := client.PostResource(endpoint, params, data)
	if err != nil {
		fmt.Printf("Error posting resource to %s: %v\n", endpoint, err)
		return
	}

	if outputFormat == "string" {
		fmt.Println(string(resp.Body()))
		return
	}

	var responseMap interface{}
	err = json.Unmarshal(resp.Body(), &responseMap)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return
	}

	switch outputFormat {
	case "table":
		var dataToDisplay interface{}
		if resource != "" {
			if resMap, ok := responseMap.(map[string]interface{}); ok {
				dataToDisplay = resMap[resource]
			} else {
				fmt.Println("Resource key specified, but response is not a map")
				return
			}
		} else {
			dataToDisplay = responseMap
		}
		if fields, ok := params["fields"]; ok {
			fieldsStr := strings.Split(fields.(string), ",")
			if err := DisplayOrderedTable(dataToDisplay, fieldsStr); err != nil {
				fmt.Println("Error displaying table:", err)
			}
		} else {
			if err := DisplayTable(dataToDisplay); err != nil {
				fmt.Println("Error displaying table:", err)
			}
		}

	case "json":
		prettyJson, err := PrintResponse(responseMap, true)
		if err != nil {
			fmt.Println("Error pretty printing JSON:", err)
			return
		}
		fmt.Println(prettyJson)

	case "csv":
		var dataToDisplay interface{}
		if resource != "" {
			if resMap, ok := responseMap.(map[string]interface{}); ok {
				dataToDisplay = resMap[resource]
			} else {
				fmt.Println("Resource key specified, but response is not a map")
				return
			}
		} else {
			dataToDisplay = responseMap
		}
		if fields, ok := params["fields"]; ok {
			fieldsStr := strings.Split(fields.(string), ",")
			csvString, err := AnyToCSVWithOrder(dataToDisplay, fieldsStr)
			if err != nil {
				fmt.Printf("Error printing CSV %v\n", err)
				return
			}
			fmt.Println(csvString)
		} else {
			csvString, err := AnyToCSV(dataToDisplay)
			if err != nil {
				fmt.Printf("Error printing CSV %v\n", err)
				return
			}
			fmt.Println(csvString)
		}

	default:
		fmt.Println("Unsupported output format:", outputFormat)
	}
}

// AnyToCSV converts any type to a CSV string
// It will ignore complex types and blank them
func AnyToCSV(data any) (string, error) {
	// Ensure the input data is a slice
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return "", errors.New("data is not a slice")
	}

	// Handle empty slice
	if v.Len() == 0 {
		return "", nil
	}

	// Check if the first element is a map
	// When firstElem is extracted, it might still be of type interface{}.
	// In such cases, use firstElem = firstElem.Elem() to get the underlying value.
	// This step is crucial because reflect.Value can wrap interface values, and you need to access the actual value inside
	firstElem := v.Index(0)
	if firstElem.Kind() == reflect.Interface || firstElem.Kind() == reflect.Ptr {
		firstElem = firstElem.Elem()
	}
	if firstElem.Kind() != reflect.Map {
		return "", errors.New("slice elements are not maps")
	}

	// Check if map keys are strings
	if firstElem.Type().Key().Kind() != reflect.String {
		return "", errors.New("map keys are not strings")
	}

	// Collect all keys for the header
	keys := make(map[string]bool)
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		if elem.Kind() == reflect.Interface || elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}
		if elem.Kind() != reflect.Map {
			continue
		}
		for _, key := range elem.MapKeys() {
			keys[key.String()] = true
		}
	}

	// Create a slice of keys
	header := make([]string, 0, len(keys))
	for key := range keys {
		header = append(header, key)
	}

	// Prepare a buffer to write the CSV data
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	// Write the header
	if err := writer.Write(header); err != nil {
		return "", err
	}

	// Write the data rows
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		if elem.Kind() == reflect.Interface || elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}
		if elem.Kind() != reflect.Map {
			continue
		}

		row := make([]string, len(header))
		for j, key := range header {
			if value := elem.MapIndex(reflect.ValueOf(key)); value.IsValid() {
				switch v := value.Interface().(type) {
				case string:
					row[j] = v
				case int:
					row[j] = strconv.Itoa(v)
				case float64:
					row[j] = strconv.FormatFloat(v, 'f', -1, 64)
				case float32:
					row[j] = strconv.FormatFloat(float64(v), 'f', -1, 32)
				default:
					row[j] = ""
				}
			} else {
				row[j] = ""
			}
		}
		if err := writer.Write(row); err != nil {
			return "", err
		}
	}

	// Flush and check for errors
	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

// AnyToCSVWithOrder converts a slice of maps to CSV with specified key order.
func AnyToCSVWithOrder(data any, keyOrder []string) (string, error) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return "", errors.New("data is not a slice")
	}

	if v.Len() == 0 {
		return "", nil
	}

	firstElem := v.Index(0)
	if firstElem.Kind() == reflect.Interface || firstElem.Kind() == reflect.Ptr {
		firstElem = firstElem.Elem()
	}
	if firstElem.Kind() != reflect.Map {
		return "", errors.New("slice elements are not maps")
	}

	if firstElem.Type().Key().Kind() != reflect.String {
		return "", errors.New("map keys are not strings")
	}

	// Use the keyOrder as the CSV header
	header := removeEmptyStrings(keyOrder)

	// Prepare a buffer to write the CSV data
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	// Write the header
	if err := writer.Write(header); err != nil {
		return "", err
	}

	// Write the data rows
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		if elem.Kind() == reflect.Interface || elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}
		if elem.Kind() != reflect.Map {
			continue
		}

		row := make([]string, len(header))
		for j, key := range header {
			if value := elem.MapIndex(reflect.ValueOf(key)); value.IsValid() {
				switch v := value.Interface().(type) {
				case string:
					row[j] = v
				case int:
					row[j] = strconv.Itoa(v)
				case float64:
					row[j] = strconv.FormatFloat(v, 'f', -1, 64)
				case float32:
					row[j] = strconv.FormatFloat(float64(v), 'f', -1, 32)
				default:
					row[j] = ""
				}
			} else {
				row[j] = ""
			}
		}
		if err := writer.Write(row); err != nil {
			return "", err
		}
	}

	// Flush and check for errors
	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

// FetchExport downloads the exported resource in the specified format
func FetchExport(client *client.Client, endpoint string, params map[string]any, mimeType string, outputFilePath string) {
	// Determine the content type based on the export format
	//contentType, err := GetContentType(exportFormat, "")
	//if err != nil {
	//	fmt.Printf("Error determining content type: %v\n", err)
	//	return
	//}

	// Fetch the resource using the ExportResource function
	resp, err := client.ExportResource(endpoint, params, mimeType)
	if err != nil {
		fmt.Printf("Error fetching export from %s: %v\n", endpoint, err)
		return
	}

	// Determine a unique file path
	uniqueFilePath := getUniqueFilePath(outputFilePath)

	// Write the response body to the specified file
	err = os.WriteFile(uniqueFilePath, resp.Body(), os.ModePerm)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", uniqueFilePath, err)
		return
	}

	fmt.Printf("File successfully downloaded to %s\n", uniqueFilePath)
}

// create a function that takes []string and removes empty strings
func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// getUniqueFilePath generates a unique file path by adding a numerical suffix if the file already exists
func getUniqueFilePath(basePath string) string {
	dir, file := filepath.Split(basePath)
	ext := filepath.Ext(file)
	baseName := strings.TrimSuffix(file, ext)

	// Check if file exists, if not, return the base path
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		return basePath
	}

	// If file exists, generate a new file name with a numerical suffix
	i := 1
	for {
		newFileName := fmt.Sprintf("%s (%d)%s", baseName, i, ext)
		newFilePath := filepath.Join(dir, newFileName)
		if _, err := os.Stat(newFilePath); os.IsNotExist(err) {
			return newFilePath
		}
		i++
	}
}
