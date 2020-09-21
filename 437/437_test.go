package _437

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 在以root为根节点的二叉树中,寻找和为sum的路径,返回这样的路径个数
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return findPath(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

// 在以node为根节点的二叉树中,寻找包含node的路径,和为sum
// 返回这样的路径个数
func findPath(node *TreeNode, num int) int {
	if node == nil {
		return 0
	}

	res := 0
	if node.Val == num {
		res += 1

	}

	res += findPath(node.Left, num-node.Val)
	res += findPath(node.Right, num-node.Val)

	return res
}

/*
      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1
*/
func Test437(t *testing.T) {
	n := pathSum(&TreeNode{10, &TreeNode{5, &TreeNode{3, &TreeNode{3, nil, nil}, &TreeNode{-2, nil, &TreeNode{1, nil, nil}}}, &TreeNode{2, nil, &TreeNode{1, nil, nil}}}, &TreeNode{-3, nil, &TreeNode{11, nil, nil}}}, 8)
	fmt.Printf("%+v", n)
}
