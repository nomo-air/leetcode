package _143

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {

}

func Test143(t *testing.T) {
	// 1->2->3->4->5
	// 1->5->2->4->3
	reorderList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}})
}
