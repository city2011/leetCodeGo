package main

import (
	"basic"
	"fmt"
)

func main() {
	fmt.Println(maxConsecutiveAnswers("TTFTTFTT", 1))
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return basic.Max(maxConsecutiveAnswers2(answerKey, k, 'T'), maxConsecutiveAnswers2(answerKey, k, 'F')) 
}

func maxConsecutiveAnswers2(answerKey string, k int, target byte) int {
	l,sum,ans := 0,0,0
	for r,_ := range answerKey {
		if answerKey[r] != target {
			sum++
		}
		for sum > k {
			if answerKey[l] != target {
				sum--
			}
			l++
		}
		ans = basic.Max(ans, r - l + 1)
	}
	return ans
}