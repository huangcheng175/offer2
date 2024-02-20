package offer2

import "math"

func maxAreaOfIsland(grid [][]int) int {
	return maxAreaOfIslandBfs(grid)
}

func maxAreaOfIslandDfs(grid [][]int) int {

	ret := math.MinInt
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 1 {
				continue
			}
			ret = max(ret, dfs(grid, i, j))
		}
	}

	if ret < 0 {
		return 0
	}

	return ret

}

func maxAreaOfIslandBfs(grid [][]int) int {

	ret := math.MinInt
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 1 {
				continue
			}
			ret = max(ret, bfs(grid, i, j))
		}
	}

	if ret < 0 {
		return 0
	}

	return ret

}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = 0
	return 1 + dfs(grid, i+1, j) + dfs(grid, i-1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1)

}

func bfs(grid [][]int, i, j int) int {
	q := make([]int, 0)
	q = append(q, i, j)
	ret := 0
	for len(q) > 0 {
		i, j := q[0], q[1]
		q = q[2:]
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] != 1 {
			continue
		}

		ret++
		grid[i][j] = 0
		q = append(q, i+1, j)
		q = append(q, i-1, j)
		q = append(q, i, j+1)
		q = append(q, i, j-1)
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
