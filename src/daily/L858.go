package main

import (
	"fmt"
	"leetCodeGo/src/basic"
)

func mirrorReflection(p int, q int) int {
	res := ""
	common := basic.Gcd(p, q)
	fmt.Println(common)
	res += fmt.Sprintf("%d", (p/common)%2)
	res += fmt.Sprintf("%d", (q/common)%2)
	fmt.Println(res)
	if res == "11" || res == "00" {
		return 1
	}
	if res == "10" {
		return 0
	}
	if res == "01" {
		return 2
	}
	return -1
}

func main() {
	mirrorReflection(2, 1)
}
