package offer2

import (
	"fmt"
	"testing"
)

type ListNode struct {
	val  int
	next *ListNode
}

func getIntersectionNode(head1, head2 *ListNode) *ListNode {
	hc1 := countList(head1)
	hc2 := countList(head2)

	var (
		delta           int
		longer, shorter *ListNode
	)
	if hc1 > hc2 {
		delta = hc1 - hc2
		longer = head1
		shorter = head2
	} else {
		delta = hc2 - hc1
		longer = head2
		shorter = head1
	}

	// longer先走n步
	for i := 0; i < delta; i++ {
		longer = longer.next
	}

	for shorter != longer {
		shorter = shorter.next
		longer = longer.next
	}

	return shorter

}

func countList(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.next
	}

	return count
}

func TestGetIntersectionNode(t *testing.T) {
	head1 := &ListNode{val: 0}
	head2 := &ListNode{val: 0}

	n1 := &ListNode{val: 1}
	n2 := &ListNode{val: 2}
	n3 := &ListNode{val: 3}
	n4 := &ListNode{val: 4}
	n5 := &ListNode{val: 5}

	head1.next = n1
	head2.next = n2
	n2.next = n3
	n4.next = n5

	n1.next = n4
	n3.next = n4

	fmt.Println(getIntersectionNode(head1, head2))

}
