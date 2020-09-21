package _104

import (
	"fmt"
	"testing"
)

/*
时间复杂度: O(n), n是树中的节点个数
空间复杂度: O(h), h是树的高度
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			firstNode := queue[0]
			queue = queue[1:]
			if firstNode.Left != nil {
				queue = append(queue, firstNode.Left)
			}
			if firstNode.Right != nil {
				queue = append(queue, firstNode.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

func Test104(t *testing.T) {
	fmt.Printf("递归：%d \n", maxDepth(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
	fmt.Printf("非递归：%d \n", maxDepth1(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
}
