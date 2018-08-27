package leetcode

import (
	"math"
)

type position struct {
	x int
	y int
}

// use BFS
func updateMatrix(matrix [][]int) [][]int {

	queue := make([]*position, 0)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, &position{x: i, y: j})
			} else {
				matrix[i][j] = math.MaxInt64
			}

		}
	}

	m, n := len(matrix), len(matrix[0])
	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for len(queue) > 0 {
		// dequeue
		current := queue[0]
		queue = queue[1:]

		for di := 0; di < len(directions); di++ {
			d := directions[di]
			dx := d[0]
			dy := d[1]

			i := current.x + dx
			j := current.y + dy

			currentValue := matrix[current.x][current.y]

			if (i < 0 || i >= m) || (j < 0 || j >= n) || matrix[i][j] <= currentValue+1 {
				continue
			}

			queue = append(queue, &position{x: i, y: j})
			matrix[i][j] = currentValue + 1
		}

	}

	return matrix
}
