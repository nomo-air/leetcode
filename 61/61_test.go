package _61

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	return &ListNode{}
}

func Test61(t *testing.T) {
	// 4->5->1->2->3->NULL
	s := rotateRight(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
