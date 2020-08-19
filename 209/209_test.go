package _209

import (
	"fmt"
	"testing"
)

func minSubArrayLen(s int, nums []int) int {
	start := 0
	end := 0
	sum := 0
	minArea := 2147483647 // int max
	n := len(nums) - 1

	for end <= n {
		sum += nums[end]
		for sum >= s {
			area := end - start + 1

			if area < minArea {
				minArea = area
			}

			sum -= nums[start]
			start++
		}
		end++
	}

	if minArea == 2147483647 {
		return 0
	}
	return minArea
}

func Test345(t *testing.T) {
	s := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}) // [4, 3]
	fmt.Printf("%+v", s)
}
