package main

import "leetCodeGo/src/basic"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func reverseList(head *basic.ListNode) *basic.ListNode {
	var pre *basic.ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
