package main

import (
	"container/list"
	"math"
)

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	for i := 0; i < m; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = math.MaxInt32
		}
	}
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	q := list.New()
	dis[0][0] = grid[0][0]
	q.PushFront([]int{0, 0})

	for q.Len() > 0 {
		cur := q.Remove(q.Front()).([]int)
		cx, cy := cur[0], cur[1]

		if dis[cx][cy] >= health {
			return false
		}

		if cx == m-1 && cy == n-1 {
			return true
		}

		for _, dir := range dirs {
			nx, ny := cx+dir[0], cy+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			cost := dis[cx][cy] + grid[nx][ny]
			if cost >= health {
				continue
			}
			if cost < dis[nx][ny] {
				dis[nx][ny] = cost
				if grid[nx][ny] == 0 {
					q.PushFront([]int{nx, ny})
				} else {
					q.PushBack([]int{nx, ny})
				}
			}
		}
	}
	return false
}
