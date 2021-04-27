package protein

import "errors"

var ErrStop error = errors.New("Stop Protein Error")
var ErrInvalidBase error = errors.New("Invalid Codon Errir")

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	lower := 0
	upper := 3

	for upper <= len(rna) {
		codon := rna[lower:upper]

		trans, error := FromCodon(codon)
		if error == nil {
			proteins = append(proteins, trans)
		} else if error == ErrStop {
			return proteins, nil
		} else {
			return proteins, error
		}

		lower = upper
		upper += 3
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
