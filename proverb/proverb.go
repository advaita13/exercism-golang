package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var proverbList []string

	if len(rhyme) < 1 {
		return proverbList
	}

	for i := 1; i < len(rhyme); i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i])
		proverbList = append(proverbList, line)
	}
	lastLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	proverbList = append(proverbList, lastLine)

	return proverbList
}
