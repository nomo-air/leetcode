package _445

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	return &ListNode{}
}

func Test445(t *testing.T) {
	// 7 -> 8 -> 0 -> 7
	s := addTwoNumbers(&ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}})
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
