package _237

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	return
}

func Test328(t *testing.T) {
	deleteNode(&ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}})
}
