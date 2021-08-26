package easy

func isPalindrome_List(head *ListNode1) bool {
	var fast, slow *ListNode1 = head, head
	//通过快慢指针找到中点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//如果fast不为空，则链表长度是奇数个
	if fast != nil {
		slow = slow.Next
	}
	//反转后半部分链表
	slow = reverseList(slow)

	fast = head
	//判断节点值是否相等
	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}
