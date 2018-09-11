package solution

import (
	"strconv"
)

func Solution(A int, B int) int {
	pattern := strconv.Itoa(A)
	target := strconv.Itoa(B)

	// frequency table for char
	table := make(map[string]int)
	for i := range pattern {
		p := string(pattern[i])
		table[p]++
	}

	startPosition := -1
	counter := len(pattern)

	for i := range target {
		t := string(target[i])

		if _, ok := table[t]; ok {
			// update start position
			if counter == len(pattern) {
				startPosition = i
			}

			counter--

			if counter == 0 {
				return startPosition
			}

			continue
		}

		// if not found reset the table
		startPosition = -1
		counter = len(pattern)
	}

	return startPosition
}
