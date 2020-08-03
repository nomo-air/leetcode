package _27

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v != val {
			nums[index] = v
			index++
		}
	}
	fmt.Printf("%+v", nums)
	return index
}

func removeElement2(nums []int, val int) int {
	endIndex := len(nums) - 1
	startIndex := 0
	for endIndex >= startIndex {
		if nums[startIndex] == val {
			nums[startIndex] = nums[endIndex]
			endIndex--
		} else {
			startIndex++
		}
	}

	fmt.Printf("%+v", nums)
	return startIndex
}

func Test27(t *testing.T) {
	n := removeElement2([]int{3, 2, 2, 3}, 3)
	fmt.Printf("%+v", n)
}
