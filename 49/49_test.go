package _49

import (
	"fmt"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	return [][]string{}
}

func Test49(t *testing.T) {
	n := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Printf("%+v", n)
}
