// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

// Given the head of a linked list, remove the nth node from the end of the list and return its head.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  Input: head = [1,2,3,4,5], n = 2
//  Output: [1,2,3,5]

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    tmp := head
    countNode := 0
    for tmp != nil{
        countNode++
        tmp = tmp.Next
    }    
    removePosition := countNode - n
    if n != 0{
        if removePosition > 0 {
            tmp = head
            for i:=0;i<removePosition-1;i++{
                tmp = tmp.Next
            }
            tmp.Next = tmp.Next.Next
        } else if countNode == 1 && n ==1 {
            head = nil
        } else if n == countNode{
            head = head.Next
        }
    }
    return head
}