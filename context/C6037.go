package main

import (
	"fmt"
	"sort"
)

func main() {
	res := largestInteger(65875)
	fmt.Println(res)
}

func largestInteger(num int) int {
	var pos []bool
	var odds []int
	var evens []int

	for num > 0 {
		x := num % 10
		num = num / 10
		if x % 2 == 0 {
			evens = append(evens, int(x))
			pos = append(pos, true)
		} else {
			odds = append(odds, int(x))
			pos = append(pos, false)
		}
	}

	for i := 0; i < len(pos) / 2; i++ {
		tmp := pos[i]
		pos[i] = pos[len(pos) - 1 - i]
		pos[len(pos) - 1 - i] = tmp
	}

	sort.Slice(odds, func(i, j int) bool {
		return odds[i] > odds[j]
	})
	sort.Slice(evens, func(i, j int) bool {
		return evens[i] > evens[j]
	})

	res := 0
	evenCur := 0
	oC := 0
	for i := 0; i < len(pos); i++ {
		if pos[i] {
			res = res * 10 + evens[evenCur]
			evenCur++
		} else {
			res = res * 10 + odds[oC]
			oC++
		}
	}
	return res
}
