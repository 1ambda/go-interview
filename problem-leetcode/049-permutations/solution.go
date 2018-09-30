package leetcode

func permute(nums []int) [][]int {

	permutations := make([][]int, 0)

	length := len(nums)
	vector := make([]int, length)
	taken := make([]bool, length)

	count := 0

	backtrack(count, length, nums, vector, taken, &permutations)

	return permutations
}

func backtrack(count int, length int, nums []int,vector []int, taken []bool, permutations *[][]int) {
	if count == length {
		// vector is being modified. need to copy
		p := make([]int, length)
		copy(p, vector)
		*permutations = append(*permutations, p)
		return
	}

	for i := 0; i < length; i++ {
		if taken[i] {
			continue
		}

		// mark
		vector[count] = nums[i]
		taken[i] = true

		backtrack(count+1, length, nums, vector, taken, permutations)

		// clean up
		taken[i] = false
	}
}
