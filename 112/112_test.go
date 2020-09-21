package _112

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
时间复杂度: O(n), n为树的节点个数
空间复杂度: O(h), h为树的高度
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func Test112(t *testing.T) {
	b := hasPathSum(&TreeNode{5, &TreeNode{4, &TreeNode{11, nil, &TreeNode{2, nil, nil}}, nil}, nil}, 22)
	fmt.Printf("%+v", b)
}
