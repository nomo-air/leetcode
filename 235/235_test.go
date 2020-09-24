package _235

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
时间复杂度: O(lgn), 其中n为树的节点个数
空间复杂度: O(h), 其中h为树的高度
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == nil || q == nil {
		return nil
	}

	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

/*
              6
		  /      \
		 2       8
       /  \     / \
      0   4    7   9
        /  \
       3    5
*/
func Test235(t *testing.T) {
	fmt.Printf("%+v", lowestCommonAncestor(&TreeNode{6, &TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}},
		&TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}},
		&TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}))
}
