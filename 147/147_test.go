package _147

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	return &ListNode{}
}

func Test147(t *testing.T) {
	s := insertionSortList(&ListNode{2, &ListNode{4, &ListNode{3, nil}}})
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
