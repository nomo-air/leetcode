package _3

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rp, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rp+1 < n && m[s[rp+1]] == 0 {
			// 不断地移动右指针
			m[s[rp+1]]++
			rp++
		}
		// 第 i 到 rp 个字符是一个极长的无重复字符子串
		ans = max(ans, rp-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Test345(t *testing.T) {
	n := lengthOfLongestSubstring("pwwkew")
	fmt.Printf("%+v", n)
}
