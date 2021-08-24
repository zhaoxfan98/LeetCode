package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	//把要删除的节点的下一个节点的值赋值给要删除的节点
	node.Val = node.Next.Val
	//删除要删除节点的下一个节点
	node.Next = node.Next.Next
}
