package main

import "strconv"

func main() {

}

func calPoints(ops []string) int {
	var mem []int
	for _,op := range ops {
		if op == "C" {
			mem = mem[:len(mem) - 1]
		} else if op == "D" {
			mem = append(mem, mem[len(mem) - 1] * 2)
		} else if op == "+" {
			mem = append(mem, mem[len(mem) - 1] + mem[len(mem) - 2])
		} else {
			if num,e := strconv.Atoi(op); e == nil {
				mem = append(mem, num)
			}
		}
	}

	res := 0
	for _, i := range mem {
		res += i
	}
	return res
}