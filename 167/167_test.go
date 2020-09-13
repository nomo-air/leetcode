package _167

import (
	"fmt"
	"testing"
)

/*
对撞指针
时间复杂度: O(n)
空间复杂度: O(1)
*/

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {

		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}

		} else if numbers[l]+numbers[r] < target {
			l++

		} else { // numbers[l] + numbers[r] > target
			r--
		}
	}
	return []int{}
}

func Test167(t *testing.T) {
	n := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("%+v", n)
}
