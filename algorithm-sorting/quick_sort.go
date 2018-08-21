package sorting

func QuickSort(arr []int, left int, right int) []int {

	if left < right {
		pivot := QuickPartition(arr, left, right)
		QuickSort(arr, left, pivot)
		QuickSort(arr, pivot+1, right)
	}

	return arr
}

// QuickPartition partitions array and its pivot index.
// (Hoare's partition version)
func QuickPartition(arr []int, left int, right int) (int) {

	// median-of-three pivot selection
	// order them to the left partition, pivot and right partition and use right as a pivot.
	mid := left + (right - left) / 2
	if arr[mid] < arr[left] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}
	if arr[right] < arr[left] {
		arr[left], arr[right] = arr[right], arr[left]
	}
	if arr[mid] < arr[right] {
		arr[mid], arr[right] = arr[right], arr[mid]
	}

	pivot := right
	i, j := left -1, right + 1

	for {
		for {
			j--
			if arr[j] <= arr[pivot] {
				break
			}
		}

		for {
			i++
			if arr[i] >= arr[pivot] {
				break
			}
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}
