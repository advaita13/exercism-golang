package dna

import "errors"

type Histogram map[byte]int

type DNA []byte

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, v := range d {
		e, ok := h[v]
		if ok {
			h[v] = e + 1
		} else {
			return nil, errors.New("contains an invalid nucleotide")
		}
	}

	return h, nil
}
