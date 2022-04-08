package main

import "fmt"

func main() {
	countPrimeSetBits(6, 10)
}

func countPrimeSetBits(left int, right int) int {
	mem := make([]bool, 21)
	mem[2] = true
	mem[3] = true
	mem[5] = true
	mem[7] = true
	mem[11] = true
	mem[13] = true
	mem[17] = true
	mem[19] = true

	res := 0
	for i := left; i <= right; i++ {
		if mem[countOne(i)] {
			res++
		}
	}
	return res
}

func countOne(num int) int {
	count := 0
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	fmt.Print(count)
	return count
}