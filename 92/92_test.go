package _92

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	return &ListNode{}
}

func reverseBetween2(head *ListNode, m int, n int) *ListNode {

	return &ListNode{}
}

func Test92(t *testing.T) {
	s := reverseBetween(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2, 4)
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
