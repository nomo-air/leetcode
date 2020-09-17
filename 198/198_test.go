package _198

import (
	"fmt"
	"testing"
)

/*
剪枝
时间复杂度: O(n^2)
空间复杂度: O(n)
*/

// memo[i] 表示考虑抢劫 nums[i...n) 所能获得的最大收益
var memo []int

func rob(nums []int) int {
	memo = make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	return tryRob(0, nums)
}

// 考虑抢劫nums[index...nums.size())这个范围的所有房子
func tryRob(index int, nums []int) int {
	if index >= len(nums) {
		return 0
	}

	if memo[index] != -1 {
		return memo[index]
	}

	res := 0
	for i := index; i < len(nums); i++ {
		res = max2(res, nums[i]+tryRob(i+2, nums))
	}

	memo[index] = res
	return res

}

// ---------------动态规划-----------------
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	memo[n-1] = nums[n-1]
	// memo[i] 表示考虑抢劫 nums[i...n-1] 所能获得的最大收益
	for i := n - 2; i >= 0; i-- {
		// memo[i]
		for j := i; j < n; j++ {
			if j+2 < n { // 防止数组越界
				memo[i] = max2(memo[i], nums[j]+memo[j+2])
			} else {
				memo[i] = max2(memo[i], nums[j])
			}
		}
	}

	return memo[0]
}

func max2(b, c int) int {
	if b >= c {
		return b
	}
	return c
}

func Test198(t *testing.T) {
	fmt.Printf("剪枝:%+v \n", rob([]int{2, 7, 9, 3, 1}))
	fmt.Printf("动态规划:%+v \n", rob1([]int{2, 7, 9, 3, 1}))
}
