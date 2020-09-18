package _0

import (
	"fmt"
	"testing"
)

/*
有一个背包，它的容量为C(Capacity)，现在有n种不同的物品，编号为0...n-1，其中每一件物品的重量为w(i)，价值为v(i)。
问可以向这个背包中盛放哪些物品，使得在不超过背包容量的基础上，物品的总价值最大。
*/

/*
剪枝

时间复杂度: O(n * C) 其中n为物品个数; C为背包容积
空间复杂度: O(n * C)
*/

var memo [][]int

func knapsack01(w, v []int, c int) int {
	if w == nil || v == nil || len(w) != len(v) {
		return 0
	}

	n := len(w)
	if n <= 0 || c <= 0 {
		return 0
	}

	memo = make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, c+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < c+1; j++ {
			memo[i][j] = -1
		}
	}

	return bestValue(w, v, n-1, c)
}

// 用 [0...index]的物品,填充容积为c的背包的最大价值
func bestValue(w, v []int, index, c int) int {

	if c <= 0 || index < 0 {
		return 0
	}

	if memo[index][c] != -1 {
		return memo[index][c]
	}

	res := bestValue(w, v, index-1, c)
	if c >= w[index] {
		res = max2(res, v[index]+bestValue(w, v, index-1, c-w[index]))
	}

	memo[index][c] = res
	return memo[index][c]
}

/**
背包问题
动态规划

时间复杂度: O(n * C) 其中n为物品个数; C为背包容积
空间复杂度: O(C), 只使用了C的额外空间
*/

func knapsack011(w []int, v []int, c int) int {
	if w == nil || v == nil || len(w) != len(v) {
		return 0
	}

	n := len(w)
	if n <= 0 || c <= 0 {
		return 0
	}

	memo := make([]int, c+1)

	// j当前背包容量，只选第一个物品，该物品能不能装进背包
	for j := 0; j <= c; j++ {
		if j >= w[0] {
			memo[j] = v[0]
		} else {
			memo[j] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := c; j >= w[i]; j-- {
			memo[j] = max2(memo[j], v[i]+memo[j-w[i]])
		}
	}

	return memo[c]
}

func max2(b, c int) int {
	if b >= c {
		return b
	}
	return c
}

func Test01(t *testing.T) {
	fmt.Printf("剪枝：%+v \n", knapsack01([]int{1, 2, 3}, []int{6, 10, 12}, 5))
	fmt.Printf("动态规划：%+v \n", knapsack011([]int{1, 2, 3}, []int{6, 10, 12}, 5))
}

/*
有一个容量为5的背包

| id    | 0 | 1  | 2
| weight| 1 | 2  | 3
| value	| 6 | 10 | 12

| - | 0 | 1 | 2  | 3  | 4  | 5
| 1 | 0 | 6 | 6  | 6  | 6  | 6
| 2 | 0 | 6 | 10 | 16 | 16 | 16
| 3 | 0 | 6 | 10 | 16 | 18 | 22
*/
