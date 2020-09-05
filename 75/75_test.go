package _75

import (
	"testing"
)

func sortColors(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	redIndex := 0
	blueIndex := len(nums) - 1

	for i := 0; i <= blueIndex; {
		switch nums[i] {
		case 0:
			nums[i] = nums[redIndex]
			nums[redIndex] = 0
			redIndex++
			i++
		case 2:
			nums[i] = nums[blueIndex]
			nums[blueIndex] = 2
			blueIndex--
		default:
			i++
		}
	}
	return
}

func Test27(t *testing.T) {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
