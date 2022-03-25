package main

import "fmt"

func main() {
	fmt.Println(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))
}

func findBall(grid [][]int) []int {
	res := make([]int, len(grid[0]))
	for i, j := 0, 0; j < len(grid[0]); j++ {
		res[j] = ballRes(grid, i, j)
	}
	return res
}

func ballRes(grid [][]int, i int, j int) int {
	gi, gj := i, j
	for ; gi < len(grid); gi++ {
		if grid[gi][gj] == 1 {
			if gj+1 == len(grid[0]) || grid[gi][gj+1] == -1 {
				return -1
			}
			gj++
		} else {
			if gj-1 == -1 || grid[gi][gj-1] == 1 {
				return -1
			}
			gj--
		}
	}
	return gj
}
