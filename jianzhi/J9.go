package main

import "container/list"

func main() {

}

type CQueue struct {
	sA []int
	sB []int
}

func Constructor9() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.sA = append(this.sA, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.sA) == 0 {
		return -1
	}
	res := this.sA[0]
	this.sA = this.sA[1:]
	return res
}

type CQueue2 struct {
	sA *list.List
	sB *list.List
}

func Constructor92() CQueue2 {
	return CQueue2{
		list.New(),
		list.New(),
	}
}

func (this *CQueue2) AppendTail(value int) {
	this.sA.PushBack(value)
}

func (this *CQueue2) DeleteHead() int {
	if this.sB.Len() == 0 {
		for this.sA.Len() > 0 {
			this.sB.PushBack(this.sA.Remove(this.sA.Back()))
		}
	}
	if this.sB.Len() > 0 {
		res := this.sB.Back().Value.(int)
		this.sB.Remove(this.sB.Back())
		return res
	}
	return -1
}
