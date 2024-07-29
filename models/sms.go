package models

import (
	"bufio"
	"os"
)

type SMSPayload struct {
	Message    string   `json:"message"`
	Recipients []string `json:"recipients"`
}

// ReadLines reads a file and returns a slice of its lines.
func readLines(filename string) ([]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Read lines into a slice
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
