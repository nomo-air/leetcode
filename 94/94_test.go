package _94

import (
	"fmt"
	"testing"
)

/*
二叉树的中序遍历
时间复杂度: O(n), n为树的节点个数
空间复杂度: O(h), h为树的高度
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var r []int

func inorderTraversal(root *TreeNode) []int {
	r = make([]int, 0)
	in(root)
	return r
}

func in(root *TreeNode) {
	if root == nil {
		return
	}
	in(root.Left)
	r = append(r, root.Val)
	in(root.Right)

}

func inorderTraversal1(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

/*
           1
         /  \
		2	 3
      /  \  / \
	 4	 5 6  7
    / \
   8  9
*/
func Test94(t *testing.T) {
	// 1,3,2
	nums := inorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}})
	for _, v := range nums {
		fmt.Printf("%+v \n", v)
	}
}
