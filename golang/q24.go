/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
    return recur(head)
}

func recur(node *ListNode) *ListNode{
    if node == nil{
        return nil
    }
    if node.Next == nil {
        return node
    }
    next := node.Next
    node.Next = recur(next.Next)
    next.Next = node
    return next
}