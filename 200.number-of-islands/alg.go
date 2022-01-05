package _00_number_of_islands

func numIslands(grid [][]byte) int {
	nr := len(grid)
	if nr == 0 {
		return 0
	}
	nc := len(grid[0])
	numIslands := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				numIslands++
				dfs(&grid, r, c)
			}
		}
	}

	return numIslands
}

func dfs(grid *[][]byte, r, c int) {
	nr := len(*grid)
	nc := len((*grid)[0])
	(*grid)[r][c] = '0'
	if r-1 >= 0 && (*grid)[r-1][c] == '1' {
		dfs(grid, r-1, c)
	}
	if r+1 < nr && (*grid)[r+1][c] == '1' {
		dfs(grid, r+1, c)
	}
	if c-1 >= 0 && (*grid)[r][c-1] == '1' {
		dfs(grid, r, c-1)
	}
	if c+1 < nc && (*grid)[r][c+1] == '1' {
		dfs(grid, r, c+1)
	}

}
