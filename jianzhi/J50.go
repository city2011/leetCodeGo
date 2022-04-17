package main

import "fmt"

func main() {
	res := firstUniqChar("kleeekjidasdf")
	fmt.Printf("%c", res)
}

func firstUniqChar2(s string) byte {
	mem := make(map[int32]bool)
	for _, c := range s {
		if _, ok := mem[c]; ok {
			mem[c] = true
		} else {
			mem[c] = false
		}
	}
	for _, c := range s {
		if !mem[c] {
			return byte(c)
		}
	}
	return byte(' ')
}

func firstUniqChar(s string) byte {
	pos := [26]int{}
	var q []int32

	for _,b := range s {
		// 转换成字符索引
		idx := b - 'a'
		if pos[idx] == 0 {
			// 1 表示出现过 1 次
			pos[idx] = 1
			q = append(q, idx)
		} else {
			// 至少出现一次，出现次数+1
			pos[idx]++
			// 重复出现过的 出队列
			for len(q) > 0 && pos[q[0]] > 1 {
				q = q[1:]
			}
		}
	}

	if len(q) > 0 {
		// 将队首的字符索引 转换成byte
		return byte('a' + q[0])
	} else {
		return ' '
	}
}