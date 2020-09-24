package _343

import (
	"fmt"
	"testing"
)

/*
剪枝、动态规划
时间复杂度: O(n^2)
空间复杂度: O(n)
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
	// memo[i]表示将数字i分割(至少分割成两部分)后得到的最大乘积
	if memo[n] != -1 {
		return memo[n]
	}

	res := -1
	for i := 1; i <= n-1; i++ { // 4可以分解成：1+?，2+?，3+?
		// i+(n-1)
		res = max3(res, i*(n-i), i*breakInteger(n-i))
	}
	memo[n] = res
	return res
}

// ---------------动态规划-----------------
func integerBreak1(n int) int {
	// memo[i]表示将数字i分割(至少分割成两部分)后得到的最大乘积
	memo = make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	for i := 2; i <= n; i++ {
		// 求解memo[i]
		for j := 1; j <= i-1; j++ {
			// 分解i， i = j+(i-j)
			memo[i] = max3(memo[i], j*(i-j), j*memo[i-j])
		}
	}
	return memo[n]
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

func Test343(t *testing.T) {
	fmt.Printf("剪枝:%+v \n", integerBreak(99))
	fmt.Printf("动态规划:%+v \n", integerBreak1(99))
}
