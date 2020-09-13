package _75

import (
	"fmt"
	"testing"
)

/*
三路快速排序的思想，对整个数组只遍历了一遍
时间复杂度: O(n)
空间复杂度: O(1)
*/

func sortColors(nums []int) {
	zero := -1       // [0...zero] == 0
	two := len(nums) // [two...n-1] == 2
	for i := 0; i < two; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		} else { // nums[i] == 0
			zero++
			nums[zero], nums[i] = nums[i], nums[zero]
			i++
		}
	}
	fmt.Printf("%+v", nums)
}

func Test27(t *testing.T) {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
