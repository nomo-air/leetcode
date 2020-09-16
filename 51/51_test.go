package _51

import (
	"fmt"
	"strings"
	"testing"
)

/*
时间复杂度: O(n^n)
空间复杂度: O(n)
*/
/*
对角线有7个（2*n-1）

横纵坐标相加：i+j
   0   1   2   3
00	01	02	03
               4
10	11	12	13
			   5
20	21	22	23
              6
30	31	32	33

-------------------------
横纵坐标相减：i-j+n-1（第一个为-3，用数组保存，希望第一个是0，所以+3）

00	01	02	03
               -3
10	11	12	13
			   -2
20	21	22	23
               -1
30	31	32	33
  3   2   1    0
*/

var col []bool  // col[i]表示第i列被占用
var dia1 []bool // dia1[i]表示第i对角线1被占用
var dia2 []bool // dia2[i]表示第i对角线被占用
var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	col = make([]bool, n)
	dia1 = make([]bool, 2*n-1)
	dia2 = make([]bool, 2*n-1)

	row := make([]int, 0)
	putQueen(n, 0, row)
	return res
}

func putQueen(n int, index int, row []int) {
	if index == n {
		res = append(res, generateBoard(n, row))
		return
	}

	for i := 0; i < n; i++ {
		// 尝试将第index行的皇后摆放在第i列
		if !col[i] && !dia1[index+i] && !dia2[index-i+n-1] {
			row = append(row, i)
			col[i] = true
			dia1[index+i] = true
			dia2[index-i+n-1] = true

			putQueen(n, index+1, row)

			col[i] = false
			dia1[index+i] = false
			dia2[index-i+n-1] = false
			row = row[:len(row)-1]
		}
	}
}

func generateBoard(n int, row []int) []string {
	board := make([]string, 0)
	for i := 0; i < n; i++ {
		charArray := make([]string, n)
		for j := 0; j < n; j++ {
			charArray[j] = "."
		}
		charArray[row[i]] = "Q"
		board = append(board, strings.Join(charArray, ""))
	}
	return board
}

func Test51(t *testing.T) {
	n := solveNQueens(8)
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n[i]); j++ {
			fmt.Printf("%+v \n", n[i][j])
		}
	}
}
