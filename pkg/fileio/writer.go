package fileio

import (
	"fmt"
	"os"
)

// WriteLines writes the provided lines to the specified file.
func WriteLines(lines []string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write line: %w", err)
		}
	}

	return nil
}
