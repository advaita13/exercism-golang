package etl

import "strings"

func Transform(old map[int][]string) map[string]int {
	new := make(map[string]int)

	for k := range old {
		for _, v := range old[k] {
			new[strings.ToLower(v)] = k
		}
	}

	return new
}
