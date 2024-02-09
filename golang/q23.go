/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    output := lists[0]
    for i, v := range lists {
        if i == 0 {
            continue
        }
        output = mergeLists(output, v)
    }
    return output
}

func mergeLists(l1, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val < l2.Val {
        l1.Next = mergeLists(l1.Next, l2)
        return l1
    } else {
        l2.Next = mergeLists(l1, l2.Next)
        return l2
    }
}