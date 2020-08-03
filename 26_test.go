package leetcode

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	startIndex := 0
	currentValue := nums[0]

	for i := 1; i <= len(nums)-1; i++ {
		if currentValue != nums[i] && startIndex != i {
			currentValue = nums[i]
			nums[startIndex+1] = currentValue
			startIndex++
		}
	}
	return startIndex + 1
}

func Test26(t *testing.T) {
	n := removeDuplicates([]int{1, 1, 2})
	fmt.Printf("%+v", n)
}
