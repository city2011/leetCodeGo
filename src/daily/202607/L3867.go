package main

import "sort"

func gcdSum(nums []int) int64 {
	prefixGcd := make([]int, len(nums))
	max := -1
	for i, num := range nums {
		if num > max {
			max = num
		}
		prefixGcd[i] = gcd(max, num)
	}

	sort.Slice(prefixGcd, func(i, j int) bool { return prefixGcd[i] < prefixGcd[j] })

	l, r, sum := 0, len(prefixGcd)-1, int64(0)
	for r > l {
		sum += int64(gcd(prefixGcd[r], prefixGcd[l]))
		r--
		l++
	}
	return sum
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
