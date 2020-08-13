package _206

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTmp
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	// 递归终止条件
	if head == nil || head.Next == nil {
		return head
	}

	retHead := reverseList2(head.Next)

	// head->next此刻指向head后面的链表的尾节点
	// head->next->next = head把head节点放在了尾部
	head.Next.Next = head
	head.Next = nil

	return retHead
}

func Test206(t *testing.T) {
	s := reverseList2(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}})
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
