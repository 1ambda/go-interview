package codility


// https://app.codility.com/programmers/lessons/5-prefix_sums/passing_cars/
func Solution(A []int) int {

	east := 0
	passing := 0

	for _, v := range A {
		if v == 0 {
			east += 1
		} else {
			passing += east
		}

		if passing > 1000000000 {
			return -1
		}
	}

	return passing
}