package _300

import (
	"fmt"
	"testing"
)

/*
剪枝
时间复杂度: O(n^2)
空间复杂度: O(n)
*/
var memo []int

func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	memo = make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}

	res := 1
	for i := 0; i < len(nums); i++ {
		res = max2(res, getMaxLength(nums, i))
	}
	return res
}

// 以 nums[index] 为结尾的最长上升子序列的长度
func getMaxLength(nums []int, index int) int {

	if memo[index] != -1 {
		return memo[index]
	}

	res := 1
	for i := 0; i <= index-1; i++ {
		if nums[index] > nums[i] {
			res = max2(res, 1+getMaxLength(nums, i))

		}
	}
	memo[index] = res
	return memo[index]
}

func max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Test300(t *testing.T) {
	fmt.Printf("剪枝：%+v \n", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
