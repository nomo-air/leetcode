package _24

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	firstNode := head
	secondNode := head.Next

	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}

func swapPairs2(head *ListNode) *ListNode {

	dummyHead := &ListNode{}
	dummyHead.Next = head

	prevNode := dummyHead

	for head != nil && head.Next != nil {
		firstNode := head
		secondNode := head.Next

		prevNode.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		prevNode = firstNode
		head = firstNode.Next
	}

	return dummyHead.Next
}

func Test2(t *testing.T) {
	// 2 -> 1 -> 4 -> 3
	s := swapPairs2(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}})
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
