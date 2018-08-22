package sorting

func MergeSortCombine(left, right []int) ([]int) {
	total := len(left) + len(right)
	combined := make([]int, 0, total)

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			combined = append(combined, right...)
			break
		}
		if len(right) == 0 {
			combined = append(combined, left...)
			break
		}

		if left[0] <= right[0] {
			combined = append(combined, left[0])
			left = left[1:]
		} else {
			combined = append(combined, right[0])
			right = right[1:]
		}
	}


	return combined
}

func MergeSort(arr []int) ([]int) {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return MergeSortCombine(left, right)
}
