package _77

import (
	"fmt"
	"testing"
)

/*
时间复杂度: O(n^k)
空间复杂度: O(k)
*/

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

var res [][]int

// 递归
func combine(n int, k int) (ans [][]int) {
	res = [][]int{}
	dfs(n, k, 1, []int{})
	return res
}

// 求解C(n,k), 当前已经找到的组合存储在path中, 需要从cur开始搜索新的元素
func dfs(n, k, cur int, path []int) {
	// temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
	if len(path)+(n-cur+1) < k {
		return
	}

	// 记录
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		res = append(res, temp)
		return
	}

	// 选择当前
	path = append(path, cur)
	dfs(n, k, cur+1, path)

	// 不选择当前
	path = path[:len(path)-1]
	dfs(n, k, cur+1, path)
	return
}

// ----------------------------------------------------------------------

// 非递归
func combine1(n int, k int) (ans [][]int) {
	res = [][]int{}
	generateCombinations(n, k, 1, []int{})
	return res
}

// 求解C(n,k), 当前已经找到的组合存储在path中, 需要从cur开始搜索新的元素
func generateCombinations(n, k, cur int, path []int) {
	// 记录
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		res = append(res, temp)
		return
	}

	// 还有k - len(path)个空位, 所以, [i...n] 中至少要有 k - len(path) 个元素
	// i最多为 n - (k - len(path)) + 1
	maxI := n - (k - len(path)) + 1
	for i := cur; i <= maxI; i++ {
		path = append(path, i)
		generateCombinations(n, k, i+1, path)
		path = path[:len(path)-1]
	}

	return
}

func Test77(t *testing.T) {
	n := combine(4, 2)
	fmt.Printf("递归：\n %+v \n \n", n)
	n = combine1(4, 2)
	fmt.Printf("非递归：\n %+v \n", n)
}
