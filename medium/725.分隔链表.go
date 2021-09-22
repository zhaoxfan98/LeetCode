package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	node := head
	for node != nil {
		length++
		node = node.Next
	}
	avg := length / k       //每个子链表平均元素的个数
	remainder := length % k //余数
	res := make([]*ListNode, k)
	for i, cur := 0, head; i < k && cur != nil; i++ {
		res[i] = cur
		pageSize := avg
		if i < remainder {
			pageSize++
		}
		//头结点已经保存
		for j := 1; j < pageSize; j++ {
			cur = cur.Next
		}
		cur, cur.Next = cur.Next, nil
	}
	return res
}
