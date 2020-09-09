package _77

import (
	"fmt"
	"testing"
)

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

// 递归实现组合型枚举
func combine(n int, k int) (ans [][]int) {
	return dfs(n, k, 1, []int{}, ans)
}

func dfs(n, k, cur int, temp []int, ans [][]int) [][]int {
	// temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
	if len(temp)+(n-cur+1) < k {
		return ans
	}

	// 记录
	if len(temp) == k {
		comb := make([]int, k)
		copy(comb, temp)
		ans = append(ans, comb)
		return ans
	}

	// 选择当前
	temp = append(temp, cur)
	ans = dfs(n, k, cur+1, temp, ans)

	// 不选择当前
	temp = temp[:len(temp)-1]
	ans = dfs(n, k, cur+1, temp, ans)
	return ans
}

// 非递归（字典序法）实现组合型枚举
func combine2(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	temp := make([]int, 0)
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}

func Test77(t *testing.T) {
	n := combine2(4, 2)
	fmt.Printf("%+v", n)
}
