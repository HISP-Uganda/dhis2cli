package utils

import (
	"os"
)

func RemoveDuplicates(numbers []string) []string {
	uniqueNumbers := make(map[string]bool)
	var result []string

	for _, number := range numbers {
		if _, exists := uniqueNumbers[number]; !exists {
			uniqueNumbers[number] = true
			result = append(result, number)
		}
	}
	return result
}

// ReadFile reads the entire content of a file and returns it as a string.
func ReadFile(filename string) (string, error) {
	// Read the entire file content
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert content to string
	return string(content), nil
}
