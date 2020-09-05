package _27

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	l := 0
	r := len(nums) - 1
	for r >= l {
		if nums[l] == val {
			nums[l] = nums[r]
			r--
		} else {
			l++
		}
	}
	return l
}
func Test27(t *testing.T) {
	n := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Printf("%+v", n)
}
