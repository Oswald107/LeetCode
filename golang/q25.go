/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    // recursion assumes k of at least 2, k=1 is just itself
    if k == 1 {
        return head
    }

    return recur(head, k)
}

func recur(node *ListNode, k int) *ListNode {
    // base case
    if node == nil {
        return nil
    }

    // confirm there are enough nodes to swap, and get the last node
    first := node
    for i := 1; i < k; i++ {
        if node.Next == nil {
            return first
        }
        node = node.Next
    }
    last := node

    // loop through it again, but switch each child to it's parent
    prev := first
    node = prev.Next
    var next *ListNode
    for i := 1; i < k; i++ {
        next = node.Next
        node.Next = prev
        prev = node
        node = next
    }

    // first becomes last, so it's child is recurred, and we return the last which is now the first
    first.Next = recur(next, k)
    return last
}