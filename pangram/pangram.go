package pangram

import (
	"fmt"
	"strings"
	"unicode"
)

var checkFlag = false

var alphabetDict = map[rune]bool{
	'A': false,
	'B': false,
	'C': false,
	'D': false,
	'E': false,
	'F': false,
	'G': false,
	'H': false,
	'I': false,
	'J': false,
	'K': false,
	'L': false,
	'M': false,
	'N': false,
	'O': false,
	'P': false,
	'Q': false,
	'R': false,
	'S': false,
	'T': false,
	'U': false,
	'V': false,
	'W': false,
	'X': false,
	'Y': false,
	'Z': false,
}

func IsPangram(input string) bool {
	str := strings.ToUpper(input)

	for _, v := range str {
		if !unicode.IsLetter(v) {
			continue
		}
		alphabetDict[v] = !checkFlag
	}

	for _, v := range alphabetDict {
		if v == checkFlag {
			checkFlag = !checkFlag
			return false
		}
	}

	checkFlag = !checkFlag
	return true
}
