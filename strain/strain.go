package strain

type Ints []int
type Strings []string
type Lists []Ints

func (list Ints) Keep(pred func(int) bool) Ints {
	var s Ints

	for _, v := range list {
		if pred(v) {
			s = append(s, v)
		}
	}

	return s
}

func (list Ints) Discard(pred func(int) bool) Ints {
	var s Ints

	for _, v := range list {
		if !pred(v) {
			s = append(s, v)
		}
	}

	return s
}

func (list Strings) Keep(pred func(string) bool) Strings {
	var s Strings

	for _, v := range list {
		if pred(v) {
			s = append(s, v)
		}
	}

	return s
}

func (list Lists) Keep(pred func([]int) bool) Lists {
	var s Lists

	for _, v := range list {
		if pred(v) {
			s = append(s, v)
		}
	}

	return s
}
