package leetcode

func subsets(nums []int) [][]int {

	powerset := make([][]int, 0)
	subset := make([]int, 0)

	start := 0

	backtrack(nums, start, subset, &powerset)

	return powerset
}

func backtrack(nums []int, start int, subset []int, powerset *[][]int) {
	// add current subset to answer
	tmp := make([]int, len(subset))
	copy(tmp, subset)
	*powerset = append(*powerset, tmp)

	for i := 0; i < len(nums); i ++ {
		subset = append(subset, nums[i])
		backtrack(nums[i+1:], start+1, subset, powerset)
		subset = subset[:len(subset)-1]
	}
}
