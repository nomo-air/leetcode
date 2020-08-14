package _21

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	prev := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return dummyHead.Next
}

func Test328(t *testing.T) {
	s := mergeTwoLists2(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}})
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
