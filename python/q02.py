# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        return recur(l1, l2, 0)

def recur(l1, l2, carry):
    if l1 is None and l2 is None:
        if carry == 1:
            return ListNode(1, None)
        return None
    
    v1 = 0
    c1 = None
    v2 = 0
    c2 = None
    
    if l1:
        v1 = l1.val
        c1 = l1.next
    if l2:
        v2 = l2.val
        c2 = l2.next
    

    val = v1 + v2 + carry
    carry_next = 0
    if val >= 10:
        val = val % 10
        carry_next = 1
    
    return ListNode(val, recur(c1, c2, carry_next))