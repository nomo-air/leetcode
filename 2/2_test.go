package _2

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

func Test328(t *testing.T) {
	// 7 -> 0 -> 8
	s := addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}})
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
