package _104

import (
	"fmt"
	"testing"
)

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

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			firstNode := queue[0]
			fmt.Printf("cap:%d \n", cap(queue))
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
	b := maxDepth(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}})
	fmt.Printf("%+v", b)
}
