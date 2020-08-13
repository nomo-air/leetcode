package _203

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := &ListNode{}
	dummyHead.Next = head

	cur := head
	prev := dummyHead

	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}

		cur = cur.Next
	}

	return dummyHead.Next
}

func Test328(t *testing.T) {
	s := removeElements(&ListNode{1, nil}, 1)
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}

	fmt.Printf("%+v \n", s)
}
