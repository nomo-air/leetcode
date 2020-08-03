package _80

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	startIndex := 0
	startCount := 0
	startValue := -1
	for i := 0; i <= len(nums)-1; i++ {
		if startValue != nums[i] {
			nums[startIndex] = nums[i]
			startCount = 1
			startValue = nums[i]
			startIndex++
		} else if startCount < 2 {
			nums[startIndex] = nums[i]
			startCount++
			startIndex++
		} else {
			startCount++
		}
	}
	fmt.Printf("%+v", nums)
	return startIndex
}

func Test80(t *testing.T) {
	n := removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
	fmt.Printf("%+v", n)
}
