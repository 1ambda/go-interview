package solution

func Solution5(A []int) bool {

	x := -1
	y := -1
	prev := A[0]

	for i := 1; i < len(A); i++ {
		if prev > A[i] {
			if x == -1 {
				x = i - 1
				y = i
			} else {
				y = i
			}
		}

		prev = A[i]
	}

	A[x], A[y] = A[y], A[x]

	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}

	return true
}

func Solution3(A []int) bool {

	left := -1
	for i := len(A) - 1; i >= 1; i-- {
		if A[i-1] > A[i] {
			for i > 0 && A[i] == A[i-1] {
				i--
			}

			left = i
			break
		}
	}

	right := -1
	for i := 0; i < len(A)-1; i++ {
		for i > 0 && A[i] == A[i-1] {
			i--
		}

		right = i
		break
	}

	if left == right || left == -1 || right == -1 {
		return false
	}

	A[left], A[right] = A[right], A[left]

	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}

	return true
}

func Solution(A []int) bool {
	if len(A) <= 1 {
		return true
	}

	swapped := false

	for i := 0; i < len(A)-1; i++ {
		if A[i+1] < A[i] {

			if swapped {
				if A[i-1] != A[i+1] {
					break
				}

			}

			A[i], A[i+1] = A[i+1], A[i]
			swapped = true
		}
	}

	for i := len(A) - 1; i > 0; i-- {
		if A[i] < A[i-1] {
			j := i - 1
			for j > 0 && A[i] < A[j] {
				j--
			}

			A[i], A[j+1] = A[j+1], A[i]
			swapped = true
			break
		}

	}

	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}

	return true
}

func Solution2(A []int) bool {
	if len(A) <= 1 {
		return true
	}

	for i := len(A) - 1; i > 0; i-- {
		if A[i] < A[i-1] {
			j := i - 1
			for j > 0 && A[i] < A[j] {
				j--
			}

			A[i], A[j+1] = A[j+1], A[i]
			break
		}

	}

	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}

	return true
}
