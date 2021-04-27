package summultiples

func SumMultiples(limit int, nums ...int) int {
	sum := 0

	for i := 1; i < limit; i++ {
		for _, v := range nums {
			if v != 0 && i%v == 0 {
				sum += i
				break
			}
		}
	}

	return sum
}
