package main

import (
	"fmt"
	"strconv"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	var res string

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			res += strconv.Itoa(node.Val)
			if node.Right == nil && node.Left == nil {
				return
			} else if node.Left == nil {
				res += "()"
				res += "("
				dfs(node.Right)
				res += ")"
			} else if node.Right == nil {
				res += "("
				dfs(node.Left)
				res += ")"
			} else {
				res += "("
				dfs(node.Left)
				res += ")"
				res += "("
				dfs(node.Right)
				res += ")"
			}
		}
	}
	dfs(root)

	return res
}

func tree2str2(root *TreeNode) string {
	var dfs func(*TreeNode) string
	dfs = func(node *TreeNode) string {
		switch {
		case root == nil:
			return ""
		case root.Left == nil && root.Right == nil:
			return strconv.Itoa(root.Val)
		case root.Right == nil:
			return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
		default:
			return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
		}
	}
	return dfs(root)
}
