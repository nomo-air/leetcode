package _136

import (
	"fmt"
	"testing"
)

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func Test136(t *testing.T) {
	fmt.Print(singleNumber([]int{4, 1, 2, 1, 2}))
}
