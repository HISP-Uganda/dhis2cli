package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(lines) > 1 {
		return lines[1:], nil
	}

	return [][]string{}, nil
}

// ReadFileToString reads the content of a file and returns it as a string.
func ReadFileToString(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file contents
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// IsValidDHIS2UID checks if the given string matches the DHIS2 UID format.
func IsValidDHIS2UID(uid string) bool {
	// Regular expression to match 11 alphanumeric characters
	re := regexp.MustCompile(`^[A-Za-z0-9]{11}$`)
	return re.MatchString(uid)
}

// GetContentType returns the MIME type based on the format and compression method.
func GetContentType(format string, compression string) (string, error) {
	format = strings.ToLower(format)
	compression = strings.ToLower(compression)

	switch format {
	case "csv":
		switch compression {
		case "gzip":
			return "application/csv+gzip", nil
		case "zip":
			return "application/csv+zip", nil
		case "none", "":
			return "text/csv", nil
		default:
			return "application/csv", fmt.Errorf("unsupported compression type: %s. Defaulting to application/csv", compression)
		}
	case "json":
		switch compression {
		case "gzip":
			return "application/json+gzip", nil
		case "zip":
			return "application/json+zip", nil
		case "none", "":
			return "application/json", nil
		default:
			return "", fmt.Errorf("unsupported compression type: %s. Defaulting to application/json", compression)
		}
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}

func GetNonDefaultFields(s interface{}) map[string]any {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	nonDefaultFields := make(map[string]any)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		// Get the default value tag
		defaultTag := fieldType.Tag.Get("default")

		// Ignore fields where the default tag is empty
		//if defaultTag == "" {
		//	continue
		//}

		// Get the json tag; if it's empty, use the field name
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = fieldType.Name
		}

		// Parse the tag value based on the field type
		var defaultValue interface{}
		switch field.Kind() {
		case reflect.String:
			defaultValue = defaultTag
		case reflect.Int:
			defaultValue, _ = strconv.Atoi(defaultTag)
		case reflect.Bool:
			defaultValue, _ = strconv.ParseBool(defaultTag)
		default:
			defaultValue = reflect.Zero(field.Type()).Interface()
		}

		// Check if the field is NOT equal to the default value
		if !reflect.DeepEqual(field.Interface(), defaultValue) {
			// Add the field and its value to the map using the json tag as the key
			nonDefaultFields[jsonTag] = field.Interface()
		}
	}
	return nonDefaultFields
}
