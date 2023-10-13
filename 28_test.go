package offer2

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatList(head *Node) (newHead *Node) {
	newHead = head
	if head == nil {
		return
	}

	node := head
	for node != nil {
		if node.Child == nil && node.Next == nil {
			node = node.Next
			continue
		}

		nodeNext := node.Next

		//  need recursion to flat child list
		childHead := flatList(node.Child)

		tail := childHead
		for tail.Next != nil {
			tail = tail.Next
		}

		node.Next = childHead
		childHead.Prev = node

		tail.Next = nodeNext
		if nodeNext != nil {
			nodeNext.Prev = tail
		}

		node.Child = nil

	}

	return newHead
}
