package leetcode

func combinationSum(candidates []int, target int) [][]int {

	solutions := make([][]int, 0)

	current := make([]int, 0)
	sum := 0

	backtrack(candidates, target, sum, current, &solutions)

	return solutions
}

func backtrack(nums []int, target int, sum int, current []int, solutions *[][]int) {
	if sum > target {
		return
	}

	if sum == target {
		tmp := make([]int, len(current))
		copy(tmp, current)
		*solutions = append(*solutions, tmp)
		return
	}

	for i := 0; i < len(nums); i ++ {
		n := nums[i]

		current = append(current, n)
		sum += n

		backtrack(nums[i:], target, sum, current, solutions)

		current = current[:len(current)-1]
		sum -= n
	}
}
