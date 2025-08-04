package main

import (
	"container/heap"
	"sort"
)

func maxEvents(events [][]int) int {
	mx := 0
	for i := 0; i < len(events); i++ {
		mx = max(mx, events[i][1])
	}
	// 以开始时间为索引初始化会议数组，已结束时间确定数组大小，因天数1对应数组索引0，所以大小为mx+1
	groups := make([][]int, mx+1)
	for _, e := range events {
		groups[e[0]] = append(groups[e[0]], e[1])
	}

	h := &hp{}
	ans := 0
	for i, v := range groups {
		// 第i天， 弹出已结束的会议
		for h.Len() > 0 && h.IntSlice[0] < i {
			heap.Pop(h)
		}
		// 第i天，放入全部可以开始参加的会议
		for _, endDay := range v {
			heap.Push(h, endDay)
		}
		// 第i天，参加结束最早的一个会议，弹出该会议
		if h.Len() > 0 {
			ans++
			heap.Pop(h)
		}
	}
	return ans
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxEventsWithFind(events [][]int) int {
	sort.Slice(events, func(i, j int) bool { return events[i][1] < events[j][1] })
	mx := events[len(events)-1][1]
	fa := make([]int, mx+2)
	for i := range fa {
		fa[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			return find(fa[x])
		}
		return x
	}

	ans := 0
	for _, e := range events {
		x := find(e[0])
		if x <= e[1] {
			ans++
			fa[x]++
		}
	}
	return ans
}
