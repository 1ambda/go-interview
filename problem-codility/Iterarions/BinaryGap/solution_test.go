package codility

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestSolution(t *testing.T) {
	assertEqual(t, Solution(1041), 5)
	assertEqual(t, Solution(32), 0)
	assertEqual(t, Solution(9), 2)
	assertEqual(t, Solution(1), 0)

	assertEqual(t, Solution(37), 2) // 100101
	assertEqual(t, Solution(41), 2) // 101001
	assertEqual(t, Solution(40), 1) // 101000
}