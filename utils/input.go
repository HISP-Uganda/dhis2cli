package utils

import (
	"encoding/csv"
	"os"
	"regexp"
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

// IsValidDHIS2UID checks if the given string matches the DHIS2 UID format.
func IsValidDHIS2UID(uid string) bool {
	// Regular expression to match 11 alphanumeric characters
	re := regexp.MustCompile(`^[A-Za-z0-9]{11}$`)
	return re.MatchString(uid)
}
