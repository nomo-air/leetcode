package _226

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
时间复杂度: O(n), n为树中节点个数
空间复杂度: O(h), h为树的高度
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := invertTree(root.Right)
	left := invertTree(root.Left)
	root.Left = right
	root.Right = left
	return root
}

func Test226(t *testing.T) {
	b := invertTree(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}})
	fmt.Printf("%+v", b)
}
