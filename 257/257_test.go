package _257

import (
	"fmt"
	"strconv"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
时间复杂度: O(n), n为树中的节点个数
空间复杂度: O(h), h为树的高度
*/
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
		return res
	}

	leftPaths := binaryTreePaths(root.Left)
	for _, s := range leftPaths {
		val := fmt.Sprintf("%d->%s", root.Val, s)
		res = append(res, val)
	}

	rightPaths := binaryTreePaths(root.Right)
	for _, s := range rightPaths {
		val := fmt.Sprintf("%d->%s", root.Val, s)
		res = append(res, val)
	}

	return res
}

func Test257(t *testing.T) {
	s := binaryTreePaths(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}})
	fmt.Printf("%+v \n", s)
}
