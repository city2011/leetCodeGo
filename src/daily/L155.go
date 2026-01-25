package main

type MinStack struct {
	minX []int
	val  []int
}

func Constructor() MinStack {
	return MinStack{
		minX: []int{},
		val:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.val = append(this.val)
	if len(this.minX) > 0 && val <= this.minX[0] {
		this.minX = append([]int{val}, this.minX...)
	}
}

func (this *MinStack) Pop() {
	x := this.val[len(this.val)-1]
	if len(this.minX) > 0 && x == this.minX[0] {
		this.minX = this.minX[1:]
	}
}

func (this *MinStack) Top() int {
	return this.val[len(this.val)-1]
}

func (this *MinStack) GetMin() int {
	return this.minX[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
