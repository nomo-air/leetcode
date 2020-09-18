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
	dummyHead := &ListNode{}
	dummyHead.Next = head

	length := 0
	for cur := dummyHead.Next; cur != nil; cur = cur.Next {
		length++
	}

	k := length - n
	cur := dummyHead
	for i := 0; i < k; i++ {
		cur = cur.Next

	}

	cur.Next = cur.Next.Next

	return dummyHead.Next
}

func Test19(t *testing.T) {
	// 1->2->3->5
	s := removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
