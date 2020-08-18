package _145

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	return next(root, []int{})
}

func next(root *TreeNode, nums []int) (retNums []int) {
	if root == nil {
		return nums
	}

	nums = next(root.Left, nums)
	nums = next(root.Right, nums)
	nums = append(nums, root.Val)
	return nums
}

func Test145(t *testing.T) {
	// 3,2,1
	nums := postorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}})
	for _, v := range nums {
		fmt.Printf("%+v \n", v)
	}
}
