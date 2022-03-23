package main

func main() {

}

func findTarget(root *TreeNode, k int) bool {
	var dfs func(*TreeNode)
	var arr []int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		arr = append(arr, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	if len(arr) < 2 {
		return false
	}
	l,r := 0, len(arr) - 1
	for l < r {
		if arr[l]+arr[r] == k {
			return true;
		} else if arr[l] + arr[r] > k {
			r--
		} else {
			l++
		}
	}
	return false
}