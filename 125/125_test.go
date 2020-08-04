package _125

import (
	"fmt"
	"strings"
	"testing"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1
	for l < r {
		for l < r && !isOk(s[l]) {
			l++
		}
		for l < r && !isOk(s[r]) {
			r--
		}
		if l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}

func isOk(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func Test125(t *testing.T) {
	b := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Printf("%+v", b)
}
