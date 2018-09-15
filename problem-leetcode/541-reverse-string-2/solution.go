package leetcode

// 1) If there are less than k characters left, reverse all of them.
// 2) If there are less than 2k but greater than or equal to k characters,
//    then reverse the first k characters and left the other as original.

// doesn't support unicode combining chars
func reverseStr(s string, k int) string {
	jump := k * 2
	runes := []rune(s)

	for i := 0; i < len(runes); i += jump {
		j := min(i + k, len(runes))
		reverse(runes[i:j])
	}

	result := string(runes) // for debug
	return result
}

func reverse(runes []rune) []rune {
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return runes
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
