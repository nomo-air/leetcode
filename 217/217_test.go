package _217

import (
	"fmt"
	"testing"
)

func containsDuplicate(nums []int) bool {
	n := len(nums)
	numMap := make(map[int]int) // value:index
	for i := 0; i < n; i++ {
		v := nums[i]
		if _, ok := numMap[v]; ok {
			return true
		}
		numMap[v] = i
	}
	return false
}

func Test219(t *testing.T) {
	b := containsDuplicate([]int{1, 2, 3, 1})
	fmt.Printf("%+v", b)
}
