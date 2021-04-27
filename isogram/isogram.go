package isogram

import "strings"

// IsIsogram checks if given string is without repeating letters
func IsIsogram(input string) bool {
	s := strings.ToUpper(input)
	repeated := map[rune]bool{}

	for _, c := range s {
		if c == ' ' || c == '-' {
			continue
		}
		if repeated[c] {
			return false
		}
		repeated[c] = true
	}
	return true
}
