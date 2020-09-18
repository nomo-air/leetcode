package _416

import (
	"fmt"
	"testing"
)

/*
剪枝
时间复杂度: O(len(nums) * O(sum(nums)))
空间复杂度: O(len(nums) * O(sum(nums)))
*/

// memo[i][c] 表示使用索引为[0...i]的这些元素,是否可以完全填充一个容量为c的背包
// -1 表示为未计算; 0 表示不可以填充; 1 表示可以填充
var memo [][]int

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			return false
		}
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	memo = make([][]int, len(nums))
	for i := range memo {
		memo[i] = make([]int, sum/2+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return tryPartition(nums, len(nums)-1, sum/2)
}

// 使用nums[0...index], 是否可以完全填充一个容量为sum的背包
func tryPartition(nums []int, index, sum int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 || index < 0 {
		return false
	}
	if memo[index][sum] != -1 {
		return memo[index][sum] == 1
	}

	if tryPartition(nums, index-1, sum) || tryPartition(nums, index-1, sum-nums[index]) {
		memo[index][sum] = 1
	} else {
		memo[index][sum] = 0
	}

	return memo[index][sum] == 1
}

// -----------------------------------------------------------

/*
动态规划
时间复杂度: O(len(nums) * O(sum(nums)))
空间复杂度: O(len(nums) * O(sum(nums)))
*/
func canPartition1(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			return false
		}
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	n := len(nums)
	c := sum / 2

	memo := make([]bool, c+1)
	for i := 0; i <= c; i++ {
		memo[i] = nums[0] == i
	}

	for i := 1; i < n; i++ {
		for j := c; j >= nums[i]; j-- {
			memo[j] = memo[j] || memo[j-nums[i]]
		}
	}
	return memo[c]
}

func Test416(t *testing.T) {
	fmt.Printf("剪枝：%+v \n", canPartition([]int{1, 5, 11, 5}))
	fmt.Printf("动态规划：%+v \n", canPartition1([]int{1, 5, 11, 5}))
}
