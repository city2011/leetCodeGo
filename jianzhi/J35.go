package main

import "fmt"

func main() {
	head2 := Node{2, nil, nil}
	head2.Random = &head2
	head := Node{1, &head2, &head2}
	newHead := copyRandomList(&head)
	fmt.Println(newHead)
}

/**
 * Definition for a Node.
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var mem map[*Node]*Node

func copyRandomList(head *Node) *Node {
	mem = make(map[*Node]*Node)
	return dfs(head)
}

func dfs(head *Node) *Node {
	if head == nil {
		return nil
	}
	if mem[head] != nil {
		return mem[head]
	}
	newHead := Node{head.Val, nil, nil}
	mem[head] = &newHead
	newHead.Next = dfs(head.Next)
	newHead.Random = dfs(head.Random)
	return &newHead
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}

func copyRandomList3(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}
