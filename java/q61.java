/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode rotateRight(ListNode head, int k) {
        if (k == 0 || head == null) {
            return head;
        }
        
        ListNode scout = head;
        int i = k;
        while (i > 0) {
            i--;
            scout = scout.next;
            if (scout == null) {
                scout = head;
                i = i % (k-i);
            }
        }
        ListNode rotateNode = head;
        while (scout.next != null) {
            scout = scout.next;
            rotateNode = rotateNode.next;
        }

        scout.next = head;
        ListNode output = rotateNode.next;
        rotateNode.next = null;
        return output;
    }
}
