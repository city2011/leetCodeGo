package main

import "fmt"

func main() {
	buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		ss := TreeNode{preorder[0], nil, nil}
		return &ss
	}
	root := preorder[0]
	split := 0
	for i, value := range inorder {
		if value == root {
			split = i
		}
	}
	fmt.Println(split)
	rootT := TreeNode{preorder[0], buildTree(preorder[1:split+1], inorder[0:split]), buildTree(preorder[split+1:], inorder[split+1:])}
	return &rootT
}
