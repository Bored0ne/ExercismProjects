// This Package adds an Abbreviate function that replaces strings with abbreviations.
package acronym

import "strings"
import "regexp"

// Abbreviate takes a string and returns a string that is abbreviated.
func Abbreviate(s string) string {
	splitStrings := strings.Split(formatString(s), " ")

	var firstLetterOfEachString []string
	for i := 0; i < len(splitStrings); i++ {
		firstLetterOfEachString = append(firstLetterOfEachString, splitStrings[i][0:1])
	}

	return strings.ToUpper(strings.Join(firstLetterOfEachString, ""))
}

// formatString takes a string and returns a string that has been sanitized.
func formatString(s string) string {
	s = regexp.MustCompile(`['"_]+`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`[-]+`).ReplaceAllString(s, " ")
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	return s
}
