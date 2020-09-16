package _509

import (
	"fmt"
	"testing"
)

// ----------- 递归 -------------
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)

}

// ----------- 剪枝 -------------
var num = 0

func fib1(n int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	return fibByMemo(n, memo)
}

func fibByMemo(n int, memo []int) int {
	num++
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if memo[n] == -1 {
		memo[n] = fibByMemo(n-1, memo) + fibByMemo(n-2, memo)
	}
	return memo[n]
}

// ----------- 动态规划 -------------
func fib2(n int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}

func Test509(t *testing.T) {
	fmt.Printf("递归：%+v \n", fib(9))
	fmt.Printf("剪枝：%+v \n", fib1(9))
	fmt.Printf("动态规划：%+v \n", fib2(9))
}
