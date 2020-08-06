package _242

import (
	"fmt"
	"testing"
)

func isAnagram(s string, t string) bool {
	return true
}

func Test242(t *testing.T) {
	b := isAnagram("anagram", "nagaram")
	fmt.Printf("%+v", b)
}
