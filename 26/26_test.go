package _26

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	start := 0
	end := 1
	numsLen := len(nums)
	for end < numsLen {
		if nums[start] != nums[end] {
			start++
			nums[start] = nums[end]
		}
		end++
	}
	return start + 1
}
func Test26(t *testing.T) {
	n := removeDuplicates([]int{1, 1, 2})
	fmt.Printf("%+v", n)
}
