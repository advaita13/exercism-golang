package romannumerals

func ToRomanNumeral(input int) string {
	n := input
	roman := ""

	if input < 1 || input > 3000 {
		return ""
	}
}
