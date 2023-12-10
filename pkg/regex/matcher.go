// Package regex contains regex matching utilities.
package regex

import (
	"fmt"
	"regexp"
)

// MatchPattern takes a regex pattern and a string and returns the matching parts of the string.
func MatchPattern(pattern string, str string) ([]string, error) {
	compiledRegex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex: %w", err)
	}

	matches := compiledRegex.FindAllString(str, -1)
	return matches, nil
}
