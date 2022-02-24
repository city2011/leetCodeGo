package main

import (
	"strconv"
	"strings"
)

func main() {
	res := complexNumberMultiply("1+-1i", "1+-1i")
	println(res)
}

func complexNumberMultiply(num1 string, num2 string) string {
	li1, ri1 := splitComplexNumber(num1)
	li2, ri2 := splitComplexNumber(num2)

	a := li1*li2 - ri1*ri2
	b := li1*ri2 + li2*ri1

	return strconv.Itoa(a) + "+" + strconv.Itoa(b) + "i"
}

func splitComplexNumber(str string) (left int, right int) {
	ss := strings.Split(str, "+")
	left, _ = strconv.Atoi(ss[0])
	right, _ = strconv.Atoi(ss[1][:len(ss[1])-1])
	return left, right
}
