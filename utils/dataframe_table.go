package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rocketlaunchr/dataframe-go"
	"os"
)

// ConvertToMapStringStringSlice converts an any type to []map[string]string.
func ConvertToMapStringStringSlice(data any) ([]map[string]string, error) {
	// Check if data is of type []any
	slice, ok := data.([]any)
	if !ok {
		return nil, errors.New("data is not of type []any")
	}

	var result []map[string]string

	for _, v := range slice {
		// Check if each element is of type map[string]any
		if m, ok := v.(map[string]any); ok {
			stringMap := make(map[string]string)
			for key, value := range m {
				if strValue, ok := value.(string); ok {
					stringMap[key] = strValue
				} else {
					continue
					// return nil, errors.New("value is not of type string")
				}
			}
			result = append(result, stringMap)
		} else {
			return nil, errors.New("element is not of type map[string]any")
		}
	}

	return result, nil
}

// CreateDataFrameFromMap creates a DataFrame from a slice of maps and returns the table representation.
func CreateDataFrameFromMap(data []map[string]string) string {
	if len(data) == 0 {
		return ""
	}

	// Extract column names from the first map
	var columns []dataframe.Series
	for key := range data[0] {
		var values []interface{}
		for _, row := range data {
			values = append(values, row[key])
		}
		columns = append(columns, dataframe.NewSeriesString(key, nil, values...))
	}

	// Create the DataFrame
	df := dataframe.NewDataFrame(columns...)

	// Return the DataFrame table representation
	return df.Table()
}

// DisplayTable prints a []map[string]any as a table using go-pretty
func DisplayTable(data any) error {
	// Check if data is of type []any
	slice, ok := data.([]any)
	if !ok {
		return errors.New("data is not of type []any")
	}

	if len(slice) == 0 {
		fmt.Println("No data to display")
		return nil
	}

	// Convert each element to map[string]any
	var dataList []map[string]any
	for _, v := range slice {
		if m, ok := v.(map[string]any); ok {
			dataList = append(dataList, m)
		} else {
			return errors.New("element is not of type map[string]any")
		}
	}

	// Create a new table writer
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Extract and set table headers from the keys of the first map
	headers := table.Row{}
	for key := range dataList[0] {
		headers = append(headers, key)
	}
	t.AppendHeader(headers)

	// Append rows to the table
	for _, row := range dataList {
		tableRow := table.Row{}
		for _, header := range headers {

			value := row[header.(string)]
			formattedValue := formatValue(value)
			tableRow = append(tableRow, formattedValue)

		}
		t.AppendRow(tableRow)
	}

	t.SetStyle(table.StyleLight)
	// t.SetAllowedRowLength(140)

	// Render the table
	t.Render()
	// t.RenderCSV()
	return nil
}

// DisplayOrderedTable displays the data in a table format with ordered keys.
func DisplayOrderedTable(data any, keyOrder []string) error {
	// Check if data is of type []any
	slice, ok := data.([]any)
	if !ok {
		return errors.New("data is not of type []any")
	}

	if len(slice) == 0 {
		fmt.Println("No data to display")
		return nil
	}

	// Convert each element to map[string]any
	var dataList []map[string]any
	for _, v := range slice {
		if m, ok := v.(map[string]any); ok {
			dataList = append(dataList, m)
		} else {
			return errors.New("element is not of type map[string]any")
		}
	}

	// Create a new table writer
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Use keyOrder as the table headers
	headers := table.Row{}
	for _, key := range keyOrder {
		headers = append(headers, key)
	}
	t.AppendHeader(headers)

	// Append rows to the table
	for _, row := range dataList {
		tableRow := table.Row{}
		for _, key := range keyOrder {
			value := row[key]
			formattedValue := formatValue(value)
			tableRow = append(tableRow, formattedValue)
		}
		t.AppendRow(tableRow)
	}

	t.SetStyle(table.StyleLight)

	// Render the table
	t.Render()
	return nil
}

// formatValue formats the value for table display
func formatValue(value any) string {
	var strValue string
	switch v := value.(type) {
	case string:
		strValue = v
	case int, float64:
		strValue = fmt.Sprintf("%v", v)
	default:
		jsonValue, _ := json.Marshal(v)
		strValue = string(jsonValue)
	}

	if len(strValue) > 60 {
		return truncateString(strValue, 60)
	}
	return strValue
}

// truncateString truncates a string to a given length and adds ellipsis
func truncateString(str string, length int) string {
	if len(str) <= length {
		return str
	}
	// Decide whether to show start and end with ellipsis or just the start
	start := str[:length/2]
	end := str[len(str)-length/2:]
	return fmt.Sprintf("%s ... %s", start, end)
}
