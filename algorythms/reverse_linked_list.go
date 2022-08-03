// https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  Input: head = [1,2,3,4,5]
//  Output: [5,4,3,2,1]

package main
import "fmt"

 type ListNode struct {
   Val int
   Next *ListNode
 }

func main(){
	var a1, a2, a3 ListNode
	a1 = ListNode{Val: 25, Next: &a2}
	a2 = ListNode{Val: 15, Next: &a3}
	a3 = ListNode{Val: 5, Next: nil}
	fmt.Println(reverseList(&a1))
}

func reverseList(head *ListNode) *ListNode {
    var res *ListNode
    
    for head != nil{
        tmp := head.Next
        head.Next = res
        res = head
        head = tmp
    }
    return res
}