package offer2

import (
	"fmt"
	"testing"
)

type ListNode struct {
	val  int
	next *ListNode
}

func geNodeInLoop(head *ListNode) *ListNode {

	if head == nil || head.next == nil {
		return nil
	}

	slow := head.next
	fast := slow.next

	for slow != nil && fast != nil {
		if slow == fast {
			return slow
		}
		slow = slow.next
		fast = fast.next
		if fast != nil {
			fast = fast.next
		}
	}

	return nil
}

func detectCycle(head *ListNode) *ListNode {
	nodeInLoop := geNodeInLoop(head)
	if nodeInLoop == nil {
		return nil
	}

	// 计算圈中的节点数量
	node := nodeInLoop
	var count = 1
	for node.next != nodeInLoop {
		count++
		node = node.next
	}

	// 快指针先走
	fast := head
	for count > 0 {
		count--
		fast = fast.next
	}

	slow := head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}

	return slow
}

func TestDetectCycle(t *testing.T) {
	head := &ListNode{
		val:  1,
		next: nil,
	}

	n1 := &ListNode{val: 1}
	n2 := &ListNode{val: 2}
	n3 := &ListNode{val: 3}
	n4 := &ListNode{val: 4}
	n5 := &ListNode{val: 5}

	head.next = n1
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	n5.next = n2

	v1 := detectCycle(head)
	fmt.Println(v1)
}
