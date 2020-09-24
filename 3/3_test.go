package _3

import (
	"fmt"
	"testing"
)

/*
滑动窗口
时间复杂度: O(len(s))
空间复杂度: O(len(charset))
*/

func lengthOfLongestSubstring(s string) int {
	freq := [256]int{}

	l := 0
	r := -1 // 滑动窗口为s[l...r]
	res := 0
	n := len(s)

	// 整个循环从 l == 0; r == -1 这个空窗口开始
	// 到l == s.size(); r == s.size()-1 这个空窗口截止
	// 在每次循环里逐渐改变窗口, 维护freq, 并记录当前窗口中是否找到了一个新的最优值
	for l < n {
		if r+1 < n && freq[s[r+1]] == 0 {
			r++
			freq[s[r]]++
		} else { // r已经到头 || freq[s[r+1]] == 1
			freq[s[l]]--
			l++
		}
		res = max2(res, r-l+1)
	}
	return res
}

func max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Test345(t *testing.T) {
	n := lengthOfLongestSubstring("pwwkew")
	fmt.Printf("%+v", n)
}
