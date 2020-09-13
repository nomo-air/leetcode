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

var r [][]int

// 递归
func combine(n int, k int) (ans [][]int) {
	r = [][]int{}
	dfs(n, k, 1, []int{})
	return r
}

// 求解C(n,k), 当前已经找到的组合存储在temp中, 需要从cur开始搜索新的元素
func dfs(n, k, cur int, temp []int) {
	// temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
	if len(temp)+(n-cur+1) < k {
		return
	}

	// 记录
	if len(temp) == k {
		comb := make([]int, k)
		copy(comb, temp)
		r = append(r, comb)
		return
	}

	// 选择当前
	temp = append(temp, cur)
	dfs(n, k, cur+1, temp)

	// 不选择当前
	temp = temp[:len(temp)-1]
	dfs(n, k, cur+1, temp)
	return
}

// ----------------------------------------------------------------------

// 非递归
func combine2(n int, k int) (ans [][]int) {
	r = [][]int{}
	generateCombinations(n, k, 1, []int{})
	return r
}

// 求解C(n,k), 当前已经找到的组合存储在temp中, 需要从cur开始搜索新的元素
func generateCombinations(n, k, cur int, temp []int) {
	// 记录
	if len(temp) == k {
		comb := make([]int, k)
		copy(comb, temp)
		r = append(r, comb)
		return
	}

	// 还有k - len(temp)个空位, 所以, [i...n] 中至少要有 k - len(temp) 个元素
	// i最多为 n - (k - len(temp)) + 1
	j := n - (k - len(temp)) + 1
	for i := cur; i <= j; i++ {
		temp = append(temp, i)
		generateCombinations(n, k, i+1, temp)
		temp = temp[:len(temp)-1]
	}

	return
}

func Test77(t *testing.T) {
	n := combine(4, 2)
	fmt.Printf("%+v \n ------------------------------- \n \n", n)
	n = combine2(4, 2)
	fmt.Printf("%+v \n ------------------------------- \n \n", n)
}
