// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = -1 // not a triangle
	Equ Kind = 3  // equilateral
	Iso Kind = 2  // isosceles
	Sca Kind = 0  // scalene
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind = NaT

	if IsValid(a) && IsValid(b) && IsValid(c) {
		if a > 0 && b > 0 && c > 0 {
			if a+b >= c && a+c >= b && b+c >= a {
				k = Sca

				if a == b && b == c {
					k = Equ
				} else {
					if a == b || a == c || b == c {
						k = Iso
					}
				}
			}
		}
	}
	return k
}

func IsValid(a float64) bool {
	return !(math.IsInf(a, 0) || math.IsNaN(a))
}
