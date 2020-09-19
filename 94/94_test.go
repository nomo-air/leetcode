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

func inorderTraversal1(root *TreeNode) []int {
	r = make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for len(stack) > 0 || cur != nil {
		// push
		// 如果左子树非空，入栈，然后继续往左子树走
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		// pop
		//左子树节点为空，出栈
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		r = append(r, cur.Val)

		// 换右子树
		cur = cur.Right
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
func Test94(t *testing.T) {
	fmt.Printf("二叉树的中序遍历： \n")
	nums := inorderTraversal(&TreeNode{1, &TreeNode{2, &TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}})
	for _, v := range nums {
		fmt.Printf("%+v \n", v)
	}

	fmt.Printf("\n非递归二叉树的中序遍历： \n")
	nums1 := inorderTraversal1(&TreeNode{1, &TreeNode{2, &TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}})
	for _, v := range nums1 {
		fmt.Printf("%+v \n", v)
	}
}
