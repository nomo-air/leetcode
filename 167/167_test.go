package _167

import (
	"fmt"
	"testing"
)

func twoSum(numbers []int, target int) []int {
	numbersMap := make(map[int]int, 0)
	for i, v := range numbers {
		numbersMap[v] = i
	}

	for i, v := range numbers {
		retValue := target - v
		if valueIndex, ok := numbersMap[retValue]; ok {
			return []int{i + 1, valueIndex + 1}
		}
	}
	return []int{}
}

func Test167(t *testing.T) {
	n := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("%+v", n)
}
