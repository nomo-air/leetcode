package _79

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

func exist(board [][]byte, word string) bool {
	m = len(board)    // 行
	n = len(board[0]) // 列

	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if searchWord(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

// 从board[startX][startY]开始, 寻找word[index...word.size())
func searchWord(board [][]byte, word string, index, startX, startY int) bool {
	fmt.Printf("index:%d，len:%d \n", index, len(word)-1)
	if index == len(word)-1 {
		return board[startX][startY] == word[index]
	}

	fmt.Printf("board[startX:%d][startY:%d] (%s) == word[index:%d] (%s) \n", startX, startY, string(board[startX][startY]), index, string(word[index]))
	if board[startX][startY] == word[index] {
		visited[startX][startY] = true
		// 从startX, startY出发,向四个方向寻
		for i := 0; i < 4; i++ {
			newX := startX + d[i][0]
			newY := startY + d[i][1]

			fmt.Printf("newY:%d newX:%d inArea:%v \n", newX, newY, inArea(newX, newY))

			if inArea(newX, newY) && !visited[newX][newY] && searchWord(board, word, index+1, newX, newY) {
				return true
			}
		}
		visited[startX][startY] = false
	}
	return false

}

func inArea(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func Test79(t *testing.T) {
	b := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	w := "ABCCED"
	fmt.Printf("%+v", exist(b, w))
}
