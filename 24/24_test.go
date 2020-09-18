package _24

import (
	"fmt"
	"testing"
)

/*
时间复杂度: O(n)
空间复杂度: O(1)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	prevNode := dummyHead
	for prevNode.Next != nil && prevNode.Next.Next != nil {
		node1 := prevNode.Next
		node2 := node1.Next

		next := node2.Next

		node2.Next = node1
		node1.Next = next

		prevNode.Next = node2
		prevNode = node1
	}

	return dummyHead.Next
}

func Test24(t *testing.T) {
	// 2 -> 1 -> 4 -> 3
	s := swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}})
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
