package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{

		Use:   "regexr [file path] [regex pattern]",
		Short: "RegexGoler is a CLI tool for regex pattern matching",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			regexPattern := args[1]
			filePath := args[0]

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
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
