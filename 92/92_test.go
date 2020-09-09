package _92

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归，双指针翻转
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	return &ListNode{}
}

// 迭代，
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
