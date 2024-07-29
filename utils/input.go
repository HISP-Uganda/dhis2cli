package utils

import (
	"encoding/csv"
	"os"
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
