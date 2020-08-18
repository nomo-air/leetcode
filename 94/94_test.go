package _94

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return in(root, []int{})
}

func in(root *TreeNode, nums []int) (retNums []int) {
	if root == nil {
		return nums
	}
	nums = in(root.Left, nums)
	nums = append(nums, root.Val)
	nums = in(root.Right, nums)
	return nums
}

func Test94(t *testing.T) {
	// 1,3,2
	nums := inorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}})
	for _, v := range nums {
		fmt.Printf("%+v \n", v)
	}
}
