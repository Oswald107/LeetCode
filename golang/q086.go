/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	left_head := &ListNode{0, nil}
	left_node := left_head
	right_head := &ListNode{0, head}
	prev := right_head
	node := head
	for node != nil {
		if node.Val < x {
			left_node.Next = node
			left_node = node
			node = node.Next
			prev.Next = node
			left_node.Next = nil
		} else {
			prev = node
			node = node.Next
		}
	}

	if left_head.Next == nil {
		return right_head.Next
	}
	left_node.Next = right_head.Next
	return left_head.Next
}