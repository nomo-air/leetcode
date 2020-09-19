package _145

import (
	"fmt"
	"testing"
)

/*
二叉树的后序遍历
时间复杂度: O(n), n为树的节点个数
空间复杂度: O(h), h为树的高度
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var r []int

func postorderTraversal(root *TreeNode) []int {
	r = make([]int, 0)
	next(root)
	return r
}

func next(root *TreeNode) {
	if root == nil {
		return
	}

	next(root.Left)
	next(root.Right)
	r = append(r, root.Val)
}

// ---------------------------------------------

/*
           1
         /  \
		2	 3
      /  \  / \
	 4	 5 6  7
    / \
   8  9
*/
func Test145(t *testing.T) {
	// 3,2,1
	nums := postorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}})
	for _, v := range nums {
		fmt.Printf("%+v \n", v)
	}
}
