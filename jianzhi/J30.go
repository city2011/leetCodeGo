package main

func main() {

}

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		[]int{},
		[]int{},
	}
}

func (minStack *MinStack) Push(x int) {
	minStack.stack = append(minStack.stack, x)
	if len(minStack.minStack) == 0 || minStack.minStack[len(minStack.minStack)-1] > x {
		minStack.minStack = append(minStack.minStack, x)
	}
}

func (minStack *MinStack) Pop() {
	tmp := minStack.stack[len(minStack.stack)-1]
	minStack.stack = minStack.stack[:len(minStack.stack)-1]
	if minStack.minStack[len(minStack.minStack)-1] == tmp {
		minStack.minStack = minStack.minStack[:len(minStack.minStack)-1]
	}
}

func (minStack *MinStack) Top() int {
	return minStack.stack[len(minStack.stack)-1]
}

func (minStack *MinStack) Min() int {
	return minStack.minStack[len(minStack.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
