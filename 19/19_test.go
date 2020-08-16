package _19

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return &ListNode{}
}

func Test19(t *testing.T) {
	// 1->2->3->5
	s := removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
