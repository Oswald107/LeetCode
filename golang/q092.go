/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	node := head
	var prev *ListNode
	for left > 1 {
		prev = node
		node = node.Next
		left--
		right--
	}

	node = recur2(node, right-left+1)

	if prev == nil {
		return node
	}

	prev.Next = node
	return head
}

func recur2(node *ListNode, k int) *ListNode {
	first := node
	prev := first
	node = prev.Next
	var next *ListNode
	for i := 1; i < k; i++ {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}
	first.Next = node
	return prev
}
