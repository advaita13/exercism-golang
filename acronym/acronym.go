package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	words := strings.Fields(s)
	abb := ""

	for _, v := range words {
		abb += string(v[0])
	}

	return strings.ToUpper(abb)
}
