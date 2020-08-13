package _83

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	valMap := make(map[int]int)
	cur := head
	pre := &ListNode{}

	for cur != nil {
		if _, ok := valMap[cur.Val]; ok {
			pre.Next = cur.Next

		} else {
			valMap[cur.Val] = 1
			pre = cur

		}
		cur = cur.Next
	}

	return head
}

func Test83(t *testing.T) {
	s := deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}})
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}
}
