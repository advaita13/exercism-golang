package darts

func Score(x float64, y float64) int {
	dist := (x * x) + (y * y)

	if dist > 100 {
		return 0
	} else if dist > 25 {
		return 1
	} else if dist > 1 {
		return 5
	} else {
		return 10
	}
}
