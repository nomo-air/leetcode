package _75

import (
	"fmt"
	"testing"
)

func sortColors(nums []int) {
	l := -1
	r := len(nums)

	for i := 0; i < r; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			r = r - 1

			tmp := nums[i]
			nums[i] = nums[r]
			nums[r] = tmp

		} else { // nums[i] == 0
			l = l + 1

			tmp := nums[l]
			nums[l] = nums[i]
			nums[i] = tmp

			i++
		}
	}

	fmt.Printf("%+v", nums)
}

func Test27(t *testing.T) {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
