package main

func main() {

}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	var queue []*Node
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		level := []int{}
		for i:=0;i<size;i++ {
			level = append(level, queue[i].Val)
			if queue[i].Children != nil {
				queue = append(queue, queue[i].Children...)
			}
		}
		res = append(res, level)
	}

	return res
}