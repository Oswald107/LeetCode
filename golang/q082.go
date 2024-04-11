/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	buffer := &ListNode{head.Val - 1, head}
	var erase *int
	two := buffer
	node := head.Next

	for node != nil {
		if erase != nil && *erase == two.Next.Val {
			two.Next = node
		} else if two.Next.Val == node.Val {
			erase = &two.Next.Val
			two.Next = node
		} else {
			two = two.Next
		}
		node = node.Next
	}

	if erase != nil && *erase == two.Next.Val {
		two.Next = node
	}

	return buffer.Next
}