package src

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	cur := head
	mem := map[*Node]int{}
	nmem := map[int]*Node{}
	idx := 0

	pre := &Node{}
	newNode := pre
	for cur != nil {
		mem[cur] = idx
		pre.Next = &Node{cur.Val, nil, nil}
		pre = pre.Next
		nmem[idx] = pre
		cur = cur.Next
		idx++
	}

	cur = head
	newNode = newNode.Next

	ncur := newNode
	for cur != nil {
		ncur.Random = nmem[mem[cur.Random]]
		cur = cur.Next
		ncur = ncur.Next
	}

	return ncur
}
