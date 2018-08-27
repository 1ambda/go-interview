package codility

func Solution(N int) int {
	if N == 0 {
		return 0
	}

	n := N
	started := false
	longest := 0
	current := 0

	for n > 0 {
		reminder := n % 2
		n = n / 2

		if reminder == 1 {
			if started {
				if current > longest {
					longest = current
				}

				current = 0
			} else {
				started = true
			}

			continue
		}

		if started == true {
			current = current + 1
		}
	}

	return longest
}
