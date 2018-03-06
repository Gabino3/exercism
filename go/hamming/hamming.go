package hamming

import (
	"fmt"
)

// Distance returns the Hamming distance between two strands of DNA
// or an error if the strands are not the same length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("Strands must be of equal length")
	}

	distance := 0

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
