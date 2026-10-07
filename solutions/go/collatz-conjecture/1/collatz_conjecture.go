package collatzconjecture

import "fmt"

// CollatzConjecture calculates the number of steps it takes to reach 1 off of a positive number
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
        return 0, fmt.Errorf("input must be positive")
    }

    counter := 0
    for ; n != 1; counter++ {
        if n%2 == 0 {
            n /= 2
        } else {
            n = 3*n + 1
        }
    }

    return counter, nil
}
