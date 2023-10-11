package offer2

import (
	"fmt"
	"testing"
)

type ListNode struct {
	val  int
	next *ListNode
}

func removeNthFromEnd(head *ListNode, k int) *ListNode {
	dummy := &ListNode{val: 1000000}
	dummy.next = head
	front := head
	back := dummy

	var i int = 1
	for i <= k {
		//fmt.Println("dfdfd......")
		front = front.next
		i++
	}

	for front != nil {
		front = front.next
		back = back.next
	}

	fmt.Println(back)
	back.next = back.next.next

	return dummy.next

}

func TestRemoveNthFromEnd(t *testing.T) {

	head := &ListNode{
		val:  0,
		next: nil,
	}
	l1 := head
	l1.next = &ListNode{val: 1, next: nil}

	l1 = l1.next
	l1.next = &ListNode{val: 2, next: nil}

	vd := removeNthFromEnd(head, 3)
	printList(vd)

	l1 = l1.next
	l1.next = &ListNode{val: 3, next: nil}

	vd = removeNthFromEnd(head, 3)
	printList(vd)
	//
	//l1 = l1.next
	//l1.next = &ListNode{val: 3, next: nil}
	//l1 = l1.next
	//l1.next = &ListNode{val: 4, next: nil}
	//
	//removeNthFromEnd(l1, 3)
	//printList(l1)

}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d, ", head.val)
		head = head.next
	}
	fmt.Println()
}
