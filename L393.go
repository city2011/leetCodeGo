package main

import "fmt"

func main() {
	res := validUtf8([]int{250, 145, 145, 145, 145})
	fmt.Println(res)
}

func validUtf8(data []int) bool {
	m1 := 1 << 7
	m2 := 1 << 6

	n := len(data)

	for idx := 0; idx < len(data); idx++ {
		nByte := 0
		mask := 1 << 7
		for mask&data[idx] > 0 {
			nByte += 1
			mask = mask >> 1
		}
		if nByte == 0 {
			continue
		}
		nByte--
		if nByte > n-idx-1 || nByte == 0 {
			return false
		}
		for i := 0; i < nByte; i++ {
			idx++
			if ((data[idx] & m1) == 0) || ((data[idx] & m2) != 0) {
				return false
			}
		}
	}

	return true
}
