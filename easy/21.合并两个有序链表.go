package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var root = ListNode{-1, nil}
	current := &root

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 == nil {
		current.Next = l2
	} else {
		current.Next = l1
	}
	return root.Next
}
