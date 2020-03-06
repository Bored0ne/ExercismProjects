// package bob replicates the responses of an angsty lackadaisical teenager.
package bob

import "regexp"

// Hey tells bob a phrase and returns his response.
func Hey(remark string) string {
	var hasCaps, _ = regexp.MatchString("[A-Z]+", remark)
	var hasLowers, _ = regexp.MatchString("[a-z]+", remark)
	var hasNumbersOrSymbols, _ = regexp.MatchString("[^\\w\\s]+", remark)
	var hasSpaces, _ = regexp.MatchString("[\\s]+", remark)
	var hasQuestionMarkAtEnd, _ = regexp.MatchString("[?][\\s]*$", remark)

	if hasCaps && !hasLowers && hasQuestionMarkAtEnd {
		return "Calm down, I know what I'm doing!"
	}
	if hasQuestionMarkAtEnd {
		return "Sure."
	}
	if hasCaps && !hasLowers {
		return "Whoa, chill out!"
	}
	if (hasSpaces && !hasCaps && !hasLowers && !hasNumbersOrSymbols) || len(remark) == 0 {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
