package easy

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func reverseList(head *ListNode1) *ListNode1 {
	//双指针 后面指针指向前面指针实现节点反转
	var pre *ListNode1
	var cur *ListNode1 = head
	for cur != nil {
		//保存cur的下一个节点
		temp := cur.Next
		//将cur反转指向pre
		cur.Next = pre
		//向后移动
		pre = cur
		cur = temp
	}
	return pre
}
