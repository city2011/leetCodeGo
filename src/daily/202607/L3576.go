package main

import (
	"fmt"
	"math"
)

func main() {
	s := "9876543210"
	queries := [][]int{{0, 9}}

	res := sumAndMultiply(s, queries)
	fmt.Println(res)

	s = "10203004"
	queries = [][]int{{0, 7}, {1, 3}, {4, 6}}
	res = sumAndMultiply(s, queries)
	fmt.Println(res)

}

// wrong solution, missing mod every times in computation
// string length is 10^5 means int is 10 ^ 10 ^ 5
func sumAndMultiply(s string, queries [][]int) []int {
	size := len(s)
	sumPre := make([]int, size+1)
	sumZero := make([]int, size+1)
	ans := make([]int, len(queries))

	origin := 0

	for i, c := range s {
		num := int(c - '0')
		sumPre[i+1] = sumPre[i] + num
		if num == 0 {
			sumZero[i+1] = sumZero[i] + 1
		} else {
			origin = origin*10 + num
			sumZero[i+1] = sumZero[i]
		}
	}

	for i, q := range queries {
		l, r := q[0], q[1]
		sum := sumPre[r+1] - sumPre[l]
		rz := sumZero[size] - sumZero[r+1]

		lz := sumZero[r+1] - sumZero[l]
		x := int(float64(origin) / math.Pow(10.0, float64(size-r-1-rz)))
		x = x % int(math.Pow(10.0, float64(r-l+1-lz)))
		ans[i] = (x * sum) % 1000000007
	}
	return ans
}
