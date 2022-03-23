package main

func main() {

}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {return false}
	n,m := len(matrix),len(matrix[0])
	x,y := n - 1, 0
	for x >= 0 && y <= m - 1 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y++
		} else {
			x--
		}
	}
	return false;
}

func lengthOfLastWord(s string) int {
	res := 0
	idx := len(s) - 1
	for i := len(s) - 1; s[i] == ' '; i-- {
		idx--
	}
	for i := idx; s[i] != ' ' && i >= 0; i-- {
		res++
	}
	return res
}