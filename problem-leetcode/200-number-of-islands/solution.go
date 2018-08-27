package leetcode

func numIslands(grid [][]byte) int {
	count := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == '1' {
				count++
				hasIsland(grid, x, y)
			}
		}
	}

	return count
}

func hasIsland(grid [][]byte, x int, y int) {
	m, n := len(grid), len(grid[0])

	if x < 0 || x >= m || y < 0 || y >= n {
		return
	}

	if grid[x][y] != '1' {
		return
	}

	// marking current and adjacent vertices as visited
	grid[x][y] = '*'
	hasIsland(grid, x + 1, y)
	hasIsland(grid, x - 1, y)
	hasIsland(grid, x, y + 1)
	hasIsland(grid, x, y - 1)
}

