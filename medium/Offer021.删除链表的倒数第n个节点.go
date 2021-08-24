package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	//首先找到要删除节点的前一个节点
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	//fast为nil，只需删除头结点即可
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}
