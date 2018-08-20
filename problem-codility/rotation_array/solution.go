package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
	if len(A) == 0 || A == nil {
		return A
	}

	rotation := K % len(A)

	length := len(A)
	rotated := make([]int, 0)

	for i := 0; i < len(A); i++ {
		position := (length - rotation + i) % length
		a := A[position]
		rotated = append(rotated, a)
	}

	return rotated
}
