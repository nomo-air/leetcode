package _144

import (
	"fmt"
	"testing"
)

/*
二叉树的前序遍历
时间复杂度: O(n), n为树的节点个数
空间复杂度: O(h), h为树的高度
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var r []int

func preorderTraversal(root *TreeNode) []int {
	r = make([]int, 0)
	pre(root)
	return r
}

func pre(root *TreeNode) {
	if root == nil {
		return
	}
	r = append(r, root.Val)
	pre(root.Left)
	pre(root.Right)
}

/*
非递归二叉树的前序遍历
时间复杂度: O(n), n为树的节点个数
空间复杂度: O(h), h为树的高度
*/

func Test328(t *testing.T) {
	// 1,2,3
	nums := preorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}})
	for _, v := range nums {
		fmt.Printf("%+v \n", v)
	}
}
