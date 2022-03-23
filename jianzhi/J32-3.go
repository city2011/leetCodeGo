package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var q []*TreeNode
	var res [][]int
	q = append(q, root)
	order := true
	for len(q) != 0 {
		n := len(q)
		var tmp []int
		for i := 0; i < n;i++ {
			if order {
				tmp = append(tmp, q[i].Val)
			} else {
				tmp = append([]int{q[i].Val}, tmp...)
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		order = !order
		q = q[n:]
		res = append(res, tmp)
	}
	return res
}