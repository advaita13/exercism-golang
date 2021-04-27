package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid checks if given string is valid as per Luhn formula
func Valid(input string) bool {
	str := strings.ReplaceAll(input, " ", "")
	if len(str) < 2 {
		return false
	}
	sum := 0

	for i := len(str) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(str[i])) {
			return false
		}

		d, _ := strconv.Atoi(string(str[i]))
		if (len(str)-i)%2 == 0 {
			d = d * 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}

	return (sum%10)%2 == 0
}
