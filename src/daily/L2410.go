package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	res := matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8})
	fmt.Println(res)

	res2 := matchPlayersAndTrainers([]int{1, 1, 1}, []int{10})
	fmt.Println(res2)
}

func matchPlayersAndTrainers(players []int, trainers []int) int {
	hp1 := &hp{}
	for _, t := range trainers {
		heap.Push(hp1, t)
	}
	hp2 := &hp{}
	for _, p := range players {
		heap.Push(hp2, p)
	}

	ans := 0
	for hp1.Len() > 0 && hp2.Len() > 0 {
		t := heap.Pop(hp1).(int)
		p := heap.Pop(hp2).(int)
		if t >= p {
			ans++
		} else {
			for t < p && hp1.Len() > 0 {
				t = heap.Pop(hp1).(int)
			}
			if t >= p {
				ans++
			}
		}
	}
	return ans
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	res := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return res
}
