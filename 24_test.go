package offer2

import (
	"fmt"
	"testing"
)

type ListNode struct {
	val  int
	next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}

	dummy := &ListNode{next: head}

	pre := dummy
	cur := head
	next := head.next

	for next != nil {
		cur.next = pre

		pre = cur
		cur = next
		next = next.next
	}

	cur.next = pre

	head.next = nil

	return cur
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d, ", head.val)
		head = head.next
	}
	fmt.Println()
}

func TestReverseList(t *testing.T) {
	head := buildList()
	v := reverseList(head)

	printList(v)
}

func buildList() *ListNode {
	head := &ListNode{val: 0}
	n1 := &ListNode{val: 1}
	n2 := &ListNode{val: 2}

	head.next = n1
	n1.next = n2

	return head
}
