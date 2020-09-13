package _283

import (
	"fmt"
	"testing"
)

/*
原地(in place)解决该问题
时间复杂度: O(n)
空间复杂度: O(1)
*/
func moveZeroes(nums []int) {

	k := 0 // nums中, [0...k)的元素均为非0元素

	// 遍历到第i个元素后,保证[0...i]中所有非0元素
	// 都按照顺序排列在[0...k)中
	// 同时, [k...i] 为 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if k != i {
				nums[k], nums[i] = nums[i], nums[k]
			}
			k++
		}
	}

	fmt.Printf("%+v", nums)
}

func Test283(t *testing.T) {
	moveZeroes([]int{0, 1, 0, 3, 12, 0, 0})
}
