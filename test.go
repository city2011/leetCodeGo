package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	left := nums[:3]
	right := nums[3:]
	fmt.Println(left)
	fmt.Println(right)
}
