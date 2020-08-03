package _283

import (
	"fmt"
	"testing"
)

func moveZeroes(nums []int) {
	nonZeroNums := make([]int, 0)
	for _, n := range nums {
		if n != 0 {
			nonZeroNums = append(nonZeroNums, n)
		}
	}

	for i := range nums {
		if i < len(nonZeroNums) {
			nums[i] = nonZeroNums[i]
		} else {
			nums[i] = 0
		}
	}
	fmt.Printf("%+v", nums)
}

func Test283(t *testing.T) {
	moveZeroes([]int{0, 1, 0, 3, 12, 0, 0})
}
