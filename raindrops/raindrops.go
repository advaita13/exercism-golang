package raindrops

import "fmt"

var sounds = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(n int) string {
	result := ""

	if n%3 == 0 {
		result += sounds[3]
	}
	if n%5 == 0 {
		result += sounds[5]
	}
	if n%7 == 0 {
		result += sounds[7]
	}

	if result == "" {
		result += fmt.Sprint(n)
	}

	return result
}
