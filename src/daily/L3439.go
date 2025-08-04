package main

import "fmt"

func main() {
	fmt.Println("maxFreeTime")

	res := maxFreeTime(99, 1, []int{7, 21, 25}, []int{13, 25, 78})
	fmt.Println(res)
}

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	mx := 0
	cur := 0
	q := initQ(k + 1)
	if startTime[0] > 0 {
		q.Push(startTime[0])
		mx += startTime[0]
		cur += startTime[0]
	}

	startTime = append(startTime, eventTime)
	for i := 1; i < len(startTime); i++ {
		dis := startTime[i] - endTime[i-1]
		if len(q.q) == k+1 {
			cur += dis
			q.Push(dis)
			left := q.Pop()
			cur -= left
			mx = max(cur, mx)
		} else {
			q.Push(dis)
			mx += dis
			cur += dis
		}
	}
	return mx
}

type queue struct {
	q []int
}

func initQ(k int) *queue {
	q := make([]int, 0, k)
	return &queue{q}
}

func (q *queue) Pop() int {
	res := q.q[0]
	q.q = q.q[1:len(q.q)]
	return res
}

func (q *queue) Push(i int) {
	q.q = append(q.q, i)
}
