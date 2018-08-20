package sorting

// Space Complexity - O(1)
// Time Complexity
// - Worst: 	O(N^2)
// - Average: 	O(N^2)
// - Best: 		O(N^2)
func SelectionSort(arr []int) []int {

	min := 0

	for i := 0; i < len(arr); i++ {
		min = i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}
