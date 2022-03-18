package hamming

import "errors"

// Distance calculate distance between two DNA
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("length does not match")
	}

	var count int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
