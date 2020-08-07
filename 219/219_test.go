package _219

import (
	"fmt"
	"testing"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	numMap := make(map[int]int) // value:index
	for i := 0; i < n; i++ {
		v := nums[i]
		if _, ok := numMap[v]; ok {
			return true
		}
		numMap[v] = i

		if len(numMap) > k {
			delete(numMap, nums[i-k])
		}
	}
	return false
}

func Test219(t *testing.T) {
	b := containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)
	fmt.Printf("%+v", b)
}
