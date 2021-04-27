package collatzconjecture

import "errors"

func CollatzConjecture(input int) (int, error) {
	if input < 1 {
		return -1, errors.New("not a positive integer")
	}

	steps := 0

	for input != 1 {
		if input%2 == 0 {
			input = input / 2
		} else {
			input = input*3 + 1
		}
		steps += 1
	}

	return steps, nil
}
