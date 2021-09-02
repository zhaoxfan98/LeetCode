package easy

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
