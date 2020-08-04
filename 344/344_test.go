package _344

import (
	"fmt"
	"testing"
)

func reverseString(s []byte) {
	l := 0
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	fmt.Printf("%s", s)
}

func Test344(t *testing.T) {
	reverseString([]byte("123"))
}
