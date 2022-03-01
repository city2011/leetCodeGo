package main

import "fmt"

func main() {
	res := convert3("PAYPALISHIRING", 3)
	fmt.Println(res)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := len(s)
	sub := 2*numRows - 2
	eachL := numRows - 1
	a1 := length / sub
	a2 := length % sub
	extr := 0
	if a2 <= numRows && a2 > 0 {
		extr = 1
	} else if a2 > numRows {
		extr = a2 - numRows + 1
	}
	col := a1*eachL + extr

	mem := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		mem[i] = make([]string, col)
	}

	x := 0
	y := 0
	down := true
	for i := 0; i < length; i++ {
		mem[x][y] = s[i : i+1]
		if x == 0 {
			down = true
		}
		if x == numRows-1 {
			down = false
		}
		if down {
			x++
		} else {
			x--
			y++
		}
	}

	var res string
	for _, row := range mem {
		for _, str := range row {
			res += str
		}
	}
	return res
}

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := len(s)
	mem := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		mem[i] = make([]string, 0)
	}
	x := 0
	y := 0
	down := true
	for i := 0; i < length; i++ {
		mem[x] = append(mem[x], s[i:i+1])
		if x == 0 {
			down = true
		}
		if x == numRows-1 {
			down = false
		}
		if down {
			x++
		} else {
			x--
			y++
		}
	}

	var res string
	for _, row := range mem {
		for _, str := range row {
			res += str
		}
	}
	return res
}

func convert3(s string, numRows int) string {
	n := len(s)
	if numRows < 2 || numRows > n {
		return s
	}

	t := 2*numRows - 2
	res := make([]byte, n)
	idx := 0
	for i := 0; i < numRows; i++ {
		for j := 0; j + i < n; j += t {
			res[idx] = s[i+j]
			if i > 0 && i < numRows-1 && j+t-i < n {
				idx++
				res[idx] = s[j+t-i]
			}
			idx++
		}
	}
	return string(res)
}
