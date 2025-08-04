package main

import "fmt"

func main() {
	ans := longestSubsequence("00101001", 1)
	fmt.Println(ans)
}

func longestSubsequence(s string, k int) int {
	n := len(s)

	length := 0
	tk := k
	for tk > 0 {
		length++
		tk = tk / 2
	}

	fmt.Println(length)

	if n < length {
		return n
	}

	ans := length
	last := s[n-length:]
	val := 0
	for i := 0; i < len(last)-1; i++ {
		if last[i] == '1' {
			val += 1
		}
		val = val * 2
	}
	if last[len(last)-1] == '1' {
		val += 1
	}

	fmt.Println(val)

	if val > k {
		ans--
		length--
	}

	for i := 0; i < n-length; i++ {
		if s[i] == '0' {
			ans++
		}
	}

	return ans
}
