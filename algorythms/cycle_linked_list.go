// https://leetcode.com/problems/linked-list-cycle/

// Given head, the head of a linked list, determine if the linked list has a cycle in it.

// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

// Return true if there is a cycle in the linked list. Otherwise, return false.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main
import "fmt"

type ListNode struct {
     Val int
	 Next *ListNode
 }

func main() {
	var a1, a2, a3, a4, a5, a6 ListNode
	a1 = ListNode{Val: 25, Next: &a2}
	a2 = ListNode{Val: 25, Next: &a3}
	a3 = ListNode{Val: 25, Next: &a4}
	a4 = ListNode{Val: 25, Next: &a5}
	a5 = ListNode{Val: 25, Next: &a6}
	a6 = ListNode{Val: 25, Next: &a1}
	fmt.Println(hasCycle(&a1))
}

func hasCycle(head *ListNode) bool {
    if head == nil{
        return false
    }
    
    a := head
    b := head
    for a!= nil && a.Next != nil{
        a = a.Next.Next
        b = b.Next
        if a == b{
            return true
        } 
    }
    return false
}