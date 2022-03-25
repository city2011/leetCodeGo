package main

import "fmt"

func main() {
	fmt.Println(reverseLeftWords3("abcdefg", 2))
}

func reverseLeftWords(s string, n int) string {
	return s[n-1:len(s)] + s[0:n]
}

func reverseLeftWords2(s string, n int) string {
	bs := make([]byte, len(s))
	idx := 0
	for i := n; i < len(s); i++ {
		bs[idx] = s[i]
		idx++
	}
	for i := 0; i < n; i++ {
		bs[idx] = s[i]
		idx++
	}
	return string(bs)
}

func reverseLeftWords3(s string, n int) string {
	bs := make([]byte, len(s))
	for i := n; i < len(s)+n; i++ {
		bs[i-n] = s[i%len(s)]
	}
	return string(bs)
}
