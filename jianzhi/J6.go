package main

func main() {
	t := ListNode{2, nil}
	a1 := ListNode{1, &t}
	a2 := ListNode{3, &a1}
	reversePrint2(&a2)
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

func reversePrint2(head *ListNode) []int {
	var res []int
	var newHead *ListNode
	cur := head
	for ;cur != nil; {
		tmp := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = tmp
	}

	for ;newHead != nil; {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}
	return res
}
