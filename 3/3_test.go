package _3

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{} // 记录无重复字串字符
	n := len(s)
	rp, ans := -1, 0 // 右指针起始-1
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1]) // 向右移滑一格，移除最左边的一个字符
		}
		for rp+1 < n && m[s[rp+1]] == 0 { // 0不存在；1存在
			m[s[rp+1]] = 1
			rp++ // 不断地移动右指针
		}
		ans = max(ans, rp-i+1) // i~rp为无重复子串
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
