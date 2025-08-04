package main

import "fmt"

func kthCharacter(k int) byte {
	var dp []int
	dp = append(dp, k-1)
	count := 0
	for i := 0; dp[i] > 0; i++ {
		dp = append(dp, dp[i]-binaryCount(dp[i]))
		count++
	}
	fmt.Println(count)
	return byte('a' + count%26)
}

func binaryCount(x int) int {
	count := 0
	for x >= 1 {
		count++
		x /= 2
	}
	if count == 0 {
		return 0
	} else {
		return 1 << (count - 1)
	}
}

func main() {
	fmt.Printf("%c\n", kthCharacter(33))
	fmt.Printf("%c\n", 'b')
}
