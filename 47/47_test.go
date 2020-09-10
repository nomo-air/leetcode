package _47

import (
	"fmt"
	"testing"
)

/*
输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

// 最终结果
var result [][]int

func permuteUnique(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
	// 清空全局数组（leetcode多次执行全局变量不会消失）
	result = [][]int{}
	// todo.把nums排序
	backtrack(nums, pathNums, used)
	return result
}

// 回溯核心
// nums: 原始列表
// pathNums: 路径上的数字
// used: 是否访问过
func backtrack(nums, pathNums []int, used []bool) {
	// 结束条件：走完了，也就是路径上的数字总数等于原始列表总数
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		// 切片底层公用数据，所以要copy
		copy(tmp, pathNums)
		// 把本次结果追加到最终结果上
		result = append(result, tmp)
		return
	}

	// 开始遍历原始数组的每个数字
	for i := 0; i < len(nums); i++ {
		// 检查是否访问过
		if !used[i] {

			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 没有访问过就选择它，然后标记成已访问过的
			used[i] = true
			// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
			pathNums = append(pathNums, nums[i])
			backtrack(nums, pathNums, used)
			// 撤销刚才的选择，也就是恢复操作
			pathNums = pathNums[:len(pathNums)-1]
			// 标记成未使用
			used[i] = false
		}
	}
}

func Test47(t *testing.T) {
	fmt.Printf("%+v", permuteUnique([]int{1, 1, 2}))
}
