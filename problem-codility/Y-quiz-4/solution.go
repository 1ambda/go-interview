package solution

import (
	"math"
)

// find shortest pattern that includes all numbers
// use sliding window to find minimum days
func Solution(A []int) int {

	// build frequency table: O(N)
	table := make(map[int]int)
	for _, v := range A {
		table[v] = 1
	}

	begin := 0
	end := 0
	counter := len(table)

	min := math.MaxInt64

	// iterate array: O(N)
	for end < len(A) {
		endLocation := A[end]

		if _, ok := table[endLocation]; ok {
			table[endLocation]--
			if table[endLocation] == 0 {
				counter--
			}
		}

		end++

		// if found all required locations, then shrink current solution
		for counter == 0 {
			if end-begin < min {
				min = end - begin
			}

			beginLocation := A[begin]

			// if begin location is not found, then shrink
			if _, ok := table[beginLocation]; ok {
				table[beginLocation]++

				if table[beginLocation] > 0 {
					counter++
				}
			}

			begin++
		}
	}

	return min
}


