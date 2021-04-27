package bob

import "strings"
import "regexp"

func Hey(remark string) string {
	var isAsking = regexp.MustCompile(`\?$`).MatchString
	var isYelling = regexp.MustCompile(`\b[A-Z]+(?:\s+[A-Z]+)*\b`).MatchString
	s := strings.TrimSpace(remark)

	switch {
	case isAsking(s) && isYelling(s):
		return "Calm down, I know what I'm doing!"
	case isAsking(s):
		return "Sure."
	case isYelling(s):
		return "Whoa, chill out!"
	case s == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
