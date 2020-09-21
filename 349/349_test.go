package _349

import (
	"fmt"
	"testing"
)

/*
时间复杂度: O(n)
空间复杂度: O(n)
*/
func intersection(nums1 []int, nums2 []int) []int {
	hash := make(map[int]bool)
	res := make([]int, 0)

	for _, v := range nums1 {
		hash[v] = true
	}

	for _, v := range nums2 {
		if hash[v] == true {
			res = append(res, v)
			hash[v] = false
		}
	}
	return res
}

func Test349(t *testing.T) {
	n := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Printf("%+v", n)
}
