package codility

import (
	"math"
)

// https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/
func Solution(X int, Y int, D int) int {
	y := float64(Y)
	x := float64(X)
	d := float64(D)
	return int(math.Ceil((y - x) / d))
}
