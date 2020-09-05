package _80

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
	count := 1
	numsLen := len(nums)

	for ; end < numsLen; end++ {
		if nums[end] == nums[end-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			start++
			nums[start] = nums[end]
		}
	}
	return start + 1
}

func Test80(t *testing.T) {
	n := removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
	fmt.Printf("%+v", n)
}
