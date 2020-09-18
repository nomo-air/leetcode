package _20

import (
	"fmt"
	"testing"
)

/*
时间复杂度: O(n)
空间复杂度: O(n)
*/

var match = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			c := stack[len(stack)-1]
			if match[s[i]] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func Test20(t *testing.T) {
	fmt.Printf("%+v", isValid("(])"))
}
