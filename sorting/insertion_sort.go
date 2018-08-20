package sorting

// Space Complexity - O(1)
// Time Complexity
// - Worst: 	O(N^2)
// - Average: 	O(N^2)
// - Best: 		O(N)
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}
