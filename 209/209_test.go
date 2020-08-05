package _209

import (
	"fmt"
	"testing"
)

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。
如果不存在符合条件的子数组，返回 0。
*/

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
