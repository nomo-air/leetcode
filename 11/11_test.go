package _11

import (
	"fmt"
	"testing"
)

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxArea := 0

	for l < r {
		area := min(height[l], height[r]) * (r - l)
		if area > maxArea {
			maxArea = area
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Test11(t *testing.T) {
	n := maxArea([]int{2, 3, 4, 5, 18, 17, 6}) // 17
	fmt.Printf("%+v", n)
}
