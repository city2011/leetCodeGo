package main

import "sort"

func maxConsecutive(bottom int, top int, special []int) int {
	if len(special) == 0 {
		return top - bottom - 1
	}
	sort.Ints(special)
	maxC := 0
	for i := 0; i < len(special) - 1; i++ {
		maxC = max(special[i + 1] - special[i] - 1, maxC)
	}
	maxC = max(special[0] - bottom - 1, maxC)
	maxC = max(top - special[len(special) - 1] - 1, maxC)
	return maxC
}

func main() {
	println(maxConsecutive(
		1,
		89,
		[]int{2,4,56,7},
	))
}