package main

import "strconv"

func main() {

}

func sumAndMultiply(n int) int64 {
	sum, x, pow := 0, 0, 1
	for ; n > 0; n /= 10 {
		mod := n % 10
		if mod > 0 {
			sum += mod
			x += pow * mod
			pow *= 10
		}
	}
	return int64(x) * int64(sum)
}

func sumAndMultiply2(n int) int64 {
	x, sum := int64(0), int64(0)
	str := strconv.Itoa(n)
	for _, c := range str {
		d := int64(c - '0')
		if d > 0 {
			sum += d
			x = x*10 + d
		}
	}
	return x * sum
}
