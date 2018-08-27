package codility

// https://app.codility.com/programmers/lessons/5-prefix_sums/genomic_range_query/
// expected worst-case time complexity is O(N+M);
// expected worst-case space complexity is O(N) (not counting the storage required for input arguments).

const FactorA = 1
const FactorC = 2
const FactorG = 3
const FactorT = 4

func Solution(S string, P []int, Q []int) []int {

	// fill `0` as the first element for prefix sum operation
	// and don't need array for `T`
	cumulativeGenA := make([]int, len(S) + 1)
	cumulativeGenC := make([]int, len(S) + 1)
	cumulativeGenG := make([]int, len(S) + 1)

	for i, rune := range S {

		a, c, g := 0, 0, 0

		if string(rune) == "A" {
			a = 1
		} else if string(rune) == "C" {
			c = 1
		} else if string(rune) == "G" {
			g = 1
		}

		cumulativeGenA[i+1] = cumulativeGenA[i] + a
		cumulativeGenC[i+1] = cumulativeGenC[i] + c
		cumulativeGenG[i+1] = cumulativeGenG[i] + g
	}

	result := make([]int, 0, len(Q))

	for i := range Q {
		p := P[i]
		q := Q[i] + 1 // since we have `zero` value

		if cumulativeGenA[q]-cumulativeGenA[p] > 0 {
			result = append(result, FactorA)
		} else if cumulativeGenC[q]-cumulativeGenC[p] > 0 {
			result = append(result, FactorC)
		} else if cumulativeGenG[q]-cumulativeGenG[p] > 0 {
			result = append(result, FactorG)
		} else {
			result = append(result,FactorT)
		}
	}

	return result
}
