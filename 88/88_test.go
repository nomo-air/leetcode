package _88

import (
	"fmt"
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// nums1 = [4,5,6,0,0,0], m = 3
	// nums2 = [1,2,3],       n = 3

	p1 := len(nums1[:m]) - 1
	p2 := len(nums2) - 1

	for i := len(nums1) - 1; i > -1; i-- {
		if p1 == -1 {
			nums1[i] = nums2[p2]
			p2--
		} else if p2 == -1 {
			nums1[i] = nums1[p1]
			p1--
		} else if nums1[p1] >= nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else if nums1[p1] < nums2[p2] {
			nums1[i] = nums2[p2]
			p2--
		}

	}

	fmt.Printf("%+v", nums1)
}

func Test80(t *testing.T) {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
