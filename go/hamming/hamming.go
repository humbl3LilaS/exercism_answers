package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a and b must be of the same length")
	}
	distance := 0
	for idx := range a {
		if a[idx] != b[idx] {
			distance += 1
		}
	}
	return distance, nil
}
