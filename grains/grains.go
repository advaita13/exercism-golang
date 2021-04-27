package grains

import (
	"errors"
	"math"
)

// Total returns the sum of powers of 2 till 2^64 which is incidentally the max value of uint64
func Total() uint64 {
	return math.MaxUint64
}

// Square calculates the number of grains for a particular square, throws error if the square number is invalid
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid number of squares")
	}

	return uint64(math.Pow(2, float64(n-1))), nil
}
