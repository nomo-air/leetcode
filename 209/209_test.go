package _209

import (
	"fmt"
	"testing"
)

/*
滑动窗口的思路
时间复杂度: O(n)
空间复杂度: O(1)
*/
func minSubArrayLen(s int, nums []int) int {
	l := 0
	r := -1 // nums[l...r]为我们的滑动窗口
	sum := 0
	n := len(nums)
	res := n + 1

	for l < n { // 窗口的左边界在数组范围内,则循环继续
		if r+1 < n && sum < s {
			r++
			sum += nums[r]
		} else { // r已经到头 或者 sum >= s
			sum -= nums[l]
			l++
		}
		if sum >= s {
			res = min2(res, r-l+1)
		}
	}
	if res == n+1 {
		return 0
	}
	return res
}

func min2(b, c int) int {
	if b <= c {
		return b
	}
	return c
}

func Test345(t *testing.T) {
	s := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}) // [4, 3]
	fmt.Printf("%+v", s)
}
