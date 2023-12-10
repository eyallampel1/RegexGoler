package fileio

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLines reads a file and returns the lines as a slice of strings.
func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read lines: %w", err)
	}

	return lines, nil
}
