package _219

import (
	"fmt"
	"testing"
)

/*
时间复杂度: O(n)
空间复杂度: O(k)
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	if nums == nil || len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}

	record := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			return true
		}
		record[nums[i]] = 1
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}

	return false
}

func Test219(t *testing.T) {
	b := containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)
	fmt.Printf("%+v", b)
}
