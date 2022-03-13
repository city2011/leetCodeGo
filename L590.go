package main

func main() {

}

func postorder(root *Node) []int {
	list := List {}
	if root == nil {
		return [] int {}
	}

	dfs1(root, &list)
	return list.mem
}

func dfs1(root *Node, list *List) {
	if root == nil {
		return
	}
	for _, child := range root.Children {
		dfs1(child, list)
	}
	list.mem = append(list.mem, root.Val)
}