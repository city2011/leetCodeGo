package main

import (
	"container/heap"
	"fmt"
)

func main() {
	res := maximumProduct([]int {0,4}, 5)
	fmt.Println(res)
}

func maximumProduct(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}

	for i := 0; i < k; i++ {
		x := heap.Pop(h).(int) + 1
		heap.Push(h, x)
	}

	res := 1
	for i := 0; i < len(nums); i++ {
		res = (res * heap.Pop(h).(int)) % 1000000007
	}
	return res
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

// 这里实现了小根堆，如果想要实现大根堆可以改为 h[i] > h[j]
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *minHeap) Swap(i, j int)  {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(x interface{})  {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	res := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return res
}
