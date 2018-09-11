package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// head: 0
// tail: 1
//


func Solution(A []int) int {
	head := 0

	headCount := 0
	tailCount := 0

	for i := range A{
		coin := A[i]

		if coin == head {
			headCount++
		} else {
			tailCount++
		}
	}

	if headCount >= tailCount {
		return tailCount
	}

	return headCount
}


