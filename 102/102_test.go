package _102

import (
	"fmt"
	"testing"
)

/*
二叉树的层序遍历
时间复杂度: O(n), n为树的节点个数
空间复杂度: O(n)
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	*TreeNode
	Level int
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return res
	}
	// []map{k:level, v:node}
	queue := make([]Node, 0)
	queue = append(queue, Node{root, 0})

	for len(queue) > 0 {
		// output
		cur := queue[0]
		queue = queue[1:]
		level := cur.Level

		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], cur.Val)

		// input
		if cur.Left != nil {
			queue = append(queue, Node{cur.Left, level + 1})
		}
		if cur.Right != nil {
			queue = append(queue, Node{cur.Right, level + 1})
		}

	}
	return res
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
func Test102(t *testing.T) {
	fmt.Printf("二叉树的层序遍历： \n")
	nums := levelOrder(&TreeNode{1, &TreeNode{2, &TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}})
	for i, v := range nums {
		fmt.Printf("%+v: %+v \n", i, v)
	}
}
