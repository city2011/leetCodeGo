package main

import "fmt"

func main() {
	root := Node{
		88,
		[]*Node{},
	}
	for i := 0; i < 5; i++ {
		x := Node{
			Val: i,
		}
		root.Children = append(root.Children, &x)
	}

	res := preorder(&root)
	fmt.Println(res)
}

type Node struct {
	Val      int
	Children []*Node
}

type List struct {
	mem []int
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	list := List{}

	dfs(root, &list)
	return list.mem
}

func dfs(root *Node, list *List) {
	if root == nil {
		return
	}
	list.mem = append(list.mem, root.Val)
	for _, each := range root.Children {
		dfs(each, list)
	}
}
