package main

import (
	"fmt"
	"os"

	"github.com/eyallampel1/RegexGoler/pkg/fileio" // import your fileio package
	"github.com/eyallampel1/RegexGoler/pkg/regex"  // import your regex package
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

			lines, err := fileio.ReadLines(filePath) // Use fileio package to read lines
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
				os.Exit(1)
			}

			for _, line := range lines {
				matches, err := regex.MatchPattern(regexPattern, line) // Use regex package to match pattern
				if err != nil {
					fmt.Fprintf(os.Stderr, "Regex error: %v\n", err)
					continue
				}

				// rest of your logic here...
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
