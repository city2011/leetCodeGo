package main

import "sort"

func findLHS(nums []int) int {
	arr := nums
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	return 0
}
