package main

import (
	"container/heap"
	"math"
)

func main() {

}

func findSafeWalk(grid [][]int, health int) bool {
	m := len(grid)
	n := len(grid[0])
	dis := make([][]int, m)
	inf := math.MaxInt
	for i:=0;i< m;i++ {
		dis[i] = make([]int, n)
		for j:=0;j<n;j++{
			dis[i][j] = inf
		}
	}

	dirs := [4][2]int{{0,1},{0,-1},{1,0},{-1,0}}


	hp := &MinHeap{}
	heap.Push(hp, Item{grid[0][0], 0, 0})
	for len(*hp) > 0 {
		cur := heap.Pop(hp).(Item)
		nx, ny := cur.x, cur.y
		if cur.x == m - 1 && cur.y == n - 1 {
			return cur.Val < health
		}

		if cur.Val > dis[nx][ny] {
			continue
		}

		dis[nx][ny] = cur.Val

		for _, dir := range dirs {
			dx, dy := cur.x + dir[0], cur.y + dir[1]
			if dx < 0 || dx >=m || dy < 0 || dy >= n || dis[dx][dy] <= grid[dx][dy] + cur.Val{
				continue
			}
			dis[dx][dy] = grid[dx][dy] + cur.Val
			heap.Push(hp, Item{dis[dx][dy], dx, dy})
		}
	}

	return false
}

type Item struct {
	Val int
	x int
	y int
}

type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}
func(h MinHeap) Less(i, j int) bool{
	return h[i].Val < h[j].Val
}

func(h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func(h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	top := old[n - 1]
	*h = old[:n - 1]
	return top
}

func(h *MinHeap) Push(item interface{}) {
	*h = append(*h, item.(Item))
}

