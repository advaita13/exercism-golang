package pythagorean

import (
	"math"
)

type Triplet [3]int

func Range(min int, max int) []Triplet {
	var result []Triplet

	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			a := float64(i)
			b := float64(j)
			k := math.Sqrt(a*a + b*b)

			if math.Trunc(k) == k && int(k) <= max {
				result = append(result, Triplet{i, j, int(k)})
			}
		}
	}

	return result
}

func Sum(n int) []Triplet {
	var result []Triplet

	for i := 1; i < int(n/3)+1; i++ {
		for j := i + 1; j < int(n/2)+1; j++ {
			k := n - i - j

			if i*i+j*j == k*k {
				result = append(result, Triplet{i, j, k})
			}
		}
	}

	return result
}
