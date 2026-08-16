package hamming

import "errors"

// Distance calculates the hamming distance between two separate strands of DNA
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("Two strands of DNA must be of same length in order to calculate their Hamming distance")
    }

    var distance int
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            distance++
        }
    }

    return distance, nil
}
