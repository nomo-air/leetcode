package _82

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := &ListNode{}
	dummyHead.Next = head

	cur := head
	prev := dummyHead
	start := head // 上上个节点

	for cur != nil {
		if cur.Val == prev.Val {
			prev = cur
		} else {

			start.Next = cur
		}
		cur = cur.Next

	}

	return dummyHead.Next
}

func Test82(t *testing.T) {
	s := deleteDuplicates(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}})
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
