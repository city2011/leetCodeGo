package main

import "fmt"
import "basic"

func main() {
	res := imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
	fmt.Println(res)
}

func imageSmoother(img [][]int) [][]int {
	n := len(img)
	m := len(img[0])

	mem := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		mem[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			mem[i][j] = img[i-1][j-1] + mem[i-1][j] + mem[i][j-1] - mem[i-1][j-1]
		}
	}

	fmt.Println(mem)

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			startx := basic.Max(1, i)
			starty := basic.Max(1, j)
			endx := basic.Min(i+2, n)
			endy := basic.Min(j+2, m)

			res[i][j] = (mem[endx][endy] + mem[startx-1][starty-1] - mem[startx-1][endy] - mem[endx][starty-1]) / ((endx - startx + 1) * (endy - starty + 1))
		}
	}
	return res
}
