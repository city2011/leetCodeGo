package main

import "fmt"

func main() {
	res := sequentialDigits(10, 1234)
	fmt.Println(res)
}

func sequentialDigits(low, high int) (ans []int) {
	x0 := 12 // 第一个窗口
	pow10 := 10
	for length := 2; x0 <= high; length++ {
		pow10 *= 10
		x := x0
		for i := length; i <= 9 && x <= high; i++ {
			if x >= low {
				ans = append(ans, x)
			}
			// 窗口向右滑动，i+1 进入窗口，i+1-length 离开窗口
			x = x*10 + i + 1 - (i+1-length)*pow10
		}
		x0 = x0*10 + length + 1
	}
	return
}

func sequentialDigits1(low int, high int) []int {
	orderNums := []int{
		12, 23, 34, 45, 56, 67, 78, 89,
		123, 234, 345, 456, 567, 678, 789,
		1234, 2345, 3456, 4567, 5678, 6789,
		12345, 23456, 34567, 45678, 56789,
		123456, 234567, 345678, 456789,
		1234567, 2345678, 3456789,
		12345678, 23456789,
		123456789}

	lidx := findFirstBig(orderNums, low)
	ridx := findFirstSmall(orderNums, high)

	if ridx < lidx || orderNums[lidx] < low || orderNums[ridx] > high {
		return []int{}
	}
	return orderNums[lidx : ridx+1]
}

func findFirstBig(input []int, x int) int {
	l, r := 0, len(input)-1
	for l < r {
		mid := l + (r-l)/2
		if input[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func findFirstSmall(input []int, x int) int {
	l, r := 0, len(input)-1
	for l < r {
		mid := l + (r-l+1)/2
		if input[mid] <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
