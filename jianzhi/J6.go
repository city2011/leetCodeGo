package main

func main() {

}

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var si []int
	for head != nil {
		si = append(si, head.Val)
		head = head.Next
	}
	for i, j := 0, len(si)-1; i < j; {
		tmp := si[i]
		si[i] = si[j]
		si[j] = tmp
	}
	return si
}
