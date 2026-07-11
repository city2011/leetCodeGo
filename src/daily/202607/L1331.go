package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	input := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	res := arrayRankTransform(input)
	fmt.Println(res)
}

func arrayRankTransform(arr []int) []int {
	sarr := append([]int{}, arr...)
	sort.Slice(sarr, func(i, j int) bool { return sarr[i] < sarr[j] })
	pre := math.MaxInt32
	rec := make(map[int]int)
	idx := 1
	for _, sentity := range sarr {
		if sentity != pre {
			rec[sentity] = idx
			idx++
			pre = sentity
		}
	}

	res := make([]int, len(arr))
	for i, entity := range arr {
		res[i] = rec[entity]
	}
	return res
}
