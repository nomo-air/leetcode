package _343

import (
	"fmt"
	"testing"
)

/*
给定一个正整数:n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。

示例 2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。
*/

// ---------------剪枝-----------------
// memo[i]表示将数字i分割(至少分割成两部分)后得到的最大乘积
var memo []int

func integerBreak(n int) int {
	memo = make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	return breakInteger(n)
}

// 将n进行分割(至少分割两部分), 可以获得的最大乘积
func breakInteger(n int) int {
	if n == 1 {
		return 1
	}
	if memo[n] != -1 {
		return memo[n]
	}

	res := -1
	for i := 1; i <= n-1; i++ {
		res = max3(res, i*(n-i), i*breakInteger(n-i))
	}
	memo[n] = res
	return res
}

func max3(a, b, c int) int {
	return max2(a, max2(b, c))
}

func max2(b, c int) int {
	if b >= c {
		return b
	}
	return c
}

// ---------------动态规划-----------------
func integerBreak1(n int) int {

	memo = make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	for i := 2; i <= n; i++ {
		// 求解memo[i]
		for j := 1; j <= i-1; j++ {
			// j+(i-j)
			memo[i] = max3(memo[i], j*(i-j), j*memo[i-j])
		}
	}
	return memo[n]
}

func Test343(t *testing.T) {
	fmt.Printf("-----------------剪枝:%+v \n", integerBreak(5))
	fmt.Printf("-----------------动态规划:%+v \n", integerBreak1(5))
}
