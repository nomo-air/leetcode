package _1

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	numbersMap := make(map[int]int, 0)
	for i, v := range nums {
		numbersMap[v] = i
	}

	for i, v := range nums {
		retValue := target - v
		if valueIndex, ok := numbersMap[retValue]; ok && valueIndex != i {
			return []int{i, valueIndex}
		}
	}
	return []int{}
}

func Test1(t *testing.T) {
	n := twoSum([]int{3, 2, 4}, 6)
	fmt.Printf("%+v", n)
}
