package _200

import (
	"fmt"
	"testing"
)

var m, n int
var visited [][]bool

var d = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func numIslands(grid [][]byte) int {
	m = len(grid)
	n = len(grid[0])

	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(grid, i, j)
				res++
			}
		}
	}
	return res
}

// 从grid[x][y]的位置开始,进行floodfill
// 保证(x,y)合法,且grid[x][y]是没有被访问过的陆地
func dfs(grid [][]byte, x, y int) {
	visited[x][y] = true
	for i := 0; i < 4; i++ {
		newX := x + d[i][0]
		newY := y + d[i][1]
		if inArea(newX, newY) && !visited[newX][newY] && grid[newX][newY] == '1' {
			dfs(grid, newX, newY)

		}
	}
	return
}

func inArea(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func Test200(t *testing.T) {
	n := numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	})
	fmt.Printf("%+v", n)
}
