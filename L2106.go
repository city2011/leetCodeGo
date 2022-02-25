package main

import "math"

func main() {
	print(maximumDifference([] int {1,3,4,6,8,9,3,1,5,7,2}))
}

func maximumDifference(nums []int) int {
	minn := math.MaxInt
	maxn := -1
	for i := 0; i < len(nums); i++ {
		minn = min(minn, nums[i])
		if nums[i] > minn {
			maxn = max(nums[i]-minn, maxn)
		}
	}
	return maxn
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
