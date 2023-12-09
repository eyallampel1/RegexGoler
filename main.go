package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file_path> <regex_pattern>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := os.Args[1]
	regexPattern := os.Args[2]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regex pattern: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	toggle := false

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringIndex(line, -1)

		last := 0
		for _, match := range matches {
			fmt.Print(line[last:match[0]])

			if toggle {
				color.New(color.FgRed).Print(line[match[0]:match[1]])
			} else {
				color.New(color.FgBlue).Print(line[match[0]:match[1]])
			}
			toggle = !toggle
			last = match[1]
		}
		fmt.Println(line[last:])
	}
}
