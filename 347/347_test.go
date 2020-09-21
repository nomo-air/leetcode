package _347

import (
	"fmt"
	"testing"
)

/*
时间复杂度: O(nlogk)
空间复杂度: O(n + k)
*/
func topKFrequent(nums []int, k int) []int {
	// 统计每个元素出现的频率
	freq := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := freq[nums[i]]; ok {
			freq[nums[i]]++
		} else {
			freq[nums[i]] = 1
		}
	}
	// 维护一个长度为k的优先队列

	// 每次从freq取出元素，和优先队列最小的比较，大于则剔除最小的，加入当前的，最后返回优先队列

	return []int{}
}

func Test349(t *testing.T) {
	n := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Printf("%+v", n)
}
