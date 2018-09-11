package leetcode

import (
	"strings"
	"math"
)

// https://leetcode.com/problems/minimum-window-substring/description/

func minWindow(str string, pattern string) string {
	if len(strings.TrimSpace(str)) == 0 || len(strings.TrimSpace(pattern)) == 0 {
		return ""
	}

	// frequency table
	table := make(map[string]int)

	for i := range pattern {
		char := string(pattern[i])
		table[char]++
	}
	counter := len(table)

	solution := ""
	solutionLen := math.MaxInt64

	begin := 0
	end := 0

	for end < len(str) {
		endChar := string(str[end])

		// if end char is in the table
		if _, ok := table[endChar]; ok {
			table[endChar]--
			if table[endChar] == 0 {
				counter--
			}
		}

		end++

		// if found a solution, then move begin to
		// 1) trim solution
		// 2) or move begin ahead to find a next solution
		for counter == 0 {

			// if the current solution is shorter
			if end-begin < solutionLen {
				solution = str[begin:end]
				solutionLen = len(solution)
			}

			beginChar := string(str[begin])

			// if begin char is NOT in the table, then trim solution
			// if begin char is in the table, then
			if _, ok := table[beginChar]; ok {
				table[beginChar]++

				// remove accumulated the same char.
				// find the most recent char and increase counter to escape begin loop
				// then we will start to find a next solution
				if table[beginChar] > 0 {
					counter++
				}
			}

			begin++

		}
	}


	return solution
}
