package _198

import (
	"fmt"
	"testing"
)

// ---------------剪枝-----------------
// memo[i] 表示考虑抢劫 nums[0...i] 所能获得的最大收益
var memo []int

func rob(nums []int) int {
	memo = make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	return tryRob(nums, len(nums)-1)

}

// 考虑抢劫nums[0...index]这个范围的所有房子
func tryRob(nums []int, index int) int {
	if index < 0 {
		return 0
	}

	if memo[index] != -1 {
		return memo[index]
	}
	// 或者当前房子放弃, 考虑[0...index-1]的所有房子
	// 或者抢劫当前的房子, 考虑[0...index-2]的所有房子
	memo[index] = max2(tryRob(nums, index-1), nums[index]+tryRob(nums, index-2))
	return memo[index]
}

func max2(b, c int) int {
	if b >= c {
		return b
	}
	return c
}

// ---------------动态规划-----------------
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	memo = make([]int, len(nums))
	memo[0] = nums[0]

	for i := 1; i < n; i++ {
		// memo[i]
		if i-2 >= 0 {
			memo[i] = max2(memo[i-1], nums[i]+memo[i-2])
		} else {
			memo[i] = max2(memo[i-1], nums[i]+0)
		}
	}
	return memo[n-1]
}

func Test198(t *testing.T) {
	fmt.Printf("-----------------剪枝:%+v \n", rob([]int{2, 7, 9, 3, 1}))
	fmt.Printf("-----------------动态规划:%+v \n", rob1([]int{2, 7, 9, 3, 1}))
}
