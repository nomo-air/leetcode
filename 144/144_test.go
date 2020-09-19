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

func pre1(root *TreeNode) []int {
	r = make([]int, 0)

	stack := make([]*TreeNode, 0)
	cur := root
	stack = append(stack, cur)

	for len(stack) > 0 {

		// pop
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r = append(r, cur.Val)

		// push
		if cur != nil && cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur != nil && cur.Left != nil {
			stack = append(stack, cur.Left)
		}

	}
	return r
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
func Test144(t *testing.T) {
	fmt.Printf("二叉树的前序遍历： \n")
	nums := preorderTraversal(&TreeNode{1, &TreeNode{2, &TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}})
	for _, v := range nums {
		fmt.Printf("%+v \n", v)
	}

	fmt.Printf("\n非递归二叉树的前序遍历： \n")
	nums1 := pre1(&TreeNode{1, &TreeNode{2, &TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}})
	for _, v := range nums1 {
		fmt.Printf("%+v \n", v)
	}
}
