package main

import (
	"container/heap"
	"sort"
)

func minimumDifference(nums []int) int64 {
	n := len(nums)
	m := n / 3

	prefix := 0
	suffix := 0
	lHeap := &hp1{}
	rHeap := &hp2{}

	mems := make([][]int, m+1)
	for i := range mems {
		mems[i] = make([]int, 2)
	}

	for i := n - m; i < n; i++ {
		suffix += nums[i]
		heap.Push(rHeap, nums[i])
	}
	mems[m][1] = suffix
	for i := n - m - 1; i >= m; i-- {
		if nums[i] > rHeap.IntSlice[0] {
			t := heap.Pop(rHeap)
			suffix += nums[i] - t.(int)
			mems[i-m][1] = suffix
			heap.Push(rHeap, nums[i])
		} else {
			mems[i-m][1] = suffix
		}
	}

	for i := 0; i < m; i++ {
		prefix += nums[i]
		heap.Push(lHeap, nums[i])
	}
	mems[0][0] = prefix
	ans := mems[0][0] - mems[0][1]
	for i := m; i < n-m; i++ {
		if nums[i] < lHeap.IntSlice[0] {
			t := heap.Pop(lHeap)
			prefix -= t.(int) - nums[i]
			mems[i-m+1][0] = prefix
			heap.Push(lHeap, nums[i])
			ans = min(ans, mems[i-m+1][0]-mems[i-m+1][1])
		} else {
			mems[i-m+1][0] = prefix
			ans = min(ans, mems[i-m][0]-mems[i-m+1][1])
		}
	}
	return int64(ans)
}

type hp1 struct {
	sort.IntSlice
}

type hp2 struct {
	sort.IntSlice
}

func (h *hp2) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp2) Pop() any {
	res := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return res
}

func (h *hp1) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp1) Pop() any {
	res := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return res
}

func (h *hp1) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func main() {
	minimumDifference([]int{7, 9, 5, 8, 1, 3})
}
