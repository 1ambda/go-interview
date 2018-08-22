package leetcode

// naive solution,
// time: O(n^2)
// space: no extra space
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	// assume each input would have exactly one solution.
	return nil
}


// faster solution: use hash tabke
// time: O(n)
// space: O(n)


func twoSumHashTable(nums []int, target int) []int {

	table := make(map[int]int) // value -> index

	for i := 0; i < len(nums); i++ {
		current := nums[i]
		pairValue := target - current

		idx, exist := table[pairValue]
		if !exist {
			table[current] = i
			continue
		}

		return []int{idx, i}
	}

	// assume each input would have exactly one solution.
	return nil
}
