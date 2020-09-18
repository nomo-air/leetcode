package _203

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
不使用虚拟头结点
时间复杂度: O(n)
空间复杂度: O(1)
*/

func removeElements(head *ListNode, val int) *ListNode {
	// 需要对头结点进行特殊处理
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return head
	}

	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			nextNode := cur.Next
			cur.Next = nextNode.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

/*
使用虚拟头结点
时间复杂度: O(n)
空间复杂度: O(1)
*/
func removeElements1(head *ListNode, val int) *ListNode {
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
	fmt.Printf("不使用虚拟头结点：\n")
	s := removeElements(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}, 6)
	for s != nil {
		fmt.Printf("%+v \n", s)
		s = s.Next
	}

	fmt.Printf("\n")
	fmt.Printf("使用虚拟头结点：\n")
	s1 := removeElements1(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}, 6)
	for s1 != nil {
		fmt.Printf("%+v \n", s1)
		s1 = s1.Next
	}
}
