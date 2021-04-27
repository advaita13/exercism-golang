package reverse

func Reverse(input string) string {
	var reverse string

	for _, v := range input {
		reverse = string(v) + reverse
	}

	return reverse
}
