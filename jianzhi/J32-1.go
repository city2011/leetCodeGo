package main

func main() {

}

func levelOrder(root *TreeNode) []int {
	var q []*TreeNode
	q = append(q, root)
	var res []int

	for len(q) != 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			res = append(res, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[n:]
	}
	return res
}
