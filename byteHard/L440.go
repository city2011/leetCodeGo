package main

import (
	"fmt"
	"leetCodeGo/src/basic" // 修改包导入路径为完整模块路径
)

func main() {
	res := findKthNumber(13, 12)
	fmt.Println(res)
}

func findKthNumber(n, k int) int {
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		if steps <= k {
			k -= steps
			cur++
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

func getSteps(cur, n int) int {
	first, last := cur, cur
	steps := 0
	for first <= n {
		steps += basic.Min(n, last) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return steps
}
