package main

import "fmt"

var gNums []int
var mX int
var val int

func countMaxOrSubsets(nums []int) int {
	gNums, mX, val = nums, 0, 0
	dfs(0, 0)
	return val
}

func dfs(pos, cur int) {
	if pos == len(gNums) {
		if cur > mX {
			mX = cur
			val = 1
		} else if cur == mX {
			val++
		}
		return
	}
	dfs(pos+1, cur)
	orRes := cur | gNums[pos]
	dfs(pos+1, orRes)
}

func main() {
	res := countMaxOrSubsets([]int{2, 2, 2})
	fmt.Println(res)
}
