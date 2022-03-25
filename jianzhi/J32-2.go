package main

func main() {

}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var q []*TreeNode
	var res [][]int
	q = append(q, root)
	for len(q) != 0 {
		n := len(q)
		var tmp []int
		for i := 0; i < n; i++ {
			tmp = append(tmp, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[n:]
		res = append(res, tmp)
	}
	return res
}
