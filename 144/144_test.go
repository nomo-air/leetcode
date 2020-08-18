package _144

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return pre(root, []int{})
}

func pre(root *TreeNode, nums []int) (retNums []int) {
	if root == nil {
		return nums
	}
	nums = append(nums, root.Val)
	nums = pre(root.Left, nums)
	nums = pre(root.Right, nums)
	return nums
}

func Test328(t *testing.T) {
	// 1,2,3
	nums := preorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}})
	for _, v := range nums {
		fmt.Printf("%+v \n", v)
	}
}
