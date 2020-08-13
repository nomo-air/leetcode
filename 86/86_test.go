package _86

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {

	return &ListNode{}
}

func Test86(t *testing.T) {
	s := partition(&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{5, nil}}}}}}, 3)
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
