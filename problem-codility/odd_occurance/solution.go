package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	if len(A) == 1 {
		return A[0]
	}

	result := 0
	for i := 0; i < len(A); i++ {
		result ^= A[i]
	}

	return result
}
