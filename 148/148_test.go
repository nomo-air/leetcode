package _148

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	return &ListNode{}
}

func Test148(t *testing.T) {
	s := sortList(&ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}})
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
