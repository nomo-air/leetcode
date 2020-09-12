package _70

import (
	"fmt"
	"testing"
)

// ---------------剪枝-----------------
var memo []int

func climbStairs(n int) int {
	memo = make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	return calcWays(n)
}

func calcWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if memo[n] == -1 {
		memo[n] = calcWays(n-1) + calcWays(n-2)
	}
	return memo[n]
}

// ---------------动态规划-----------------
func climbStairs1(n int) int {
	memo = make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	memo[0] = 1
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

func Test70(t *testing.T) {
	fmt.Printf("剪枝:%+v \n", climbStairs(3))
	fmt.Printf("动态规划:%+v \n", climbStairs1(3))
}
