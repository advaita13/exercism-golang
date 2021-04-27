package accumulate

type Operation func(string) string

func Accumulate(given []string, converter Operation) []string {
	for i, v := range given {
		given[i] = converter(v)
	}

	return given
}
