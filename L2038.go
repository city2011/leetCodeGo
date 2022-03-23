package main

import "fmt"

func main() {
	res2 := winnerOfGame2("AAAAABBBBBBAAAAA")
	res := winnerOfGame("BBBAAAABB")
	fmt.Println(res, res2)
}
func winnerOfGame2(colors string) bool {
	count := make([]int, 2)
	cur := 'A'
	cnt := 0
	for _,c := range colors {
		if c != cur {
			if cnt > 2 {
				count[cur - 'A'] += cnt - 2
			}
			cur, cnt = c, 1
		} else {
			cnt++
		}
	}
	if cnt > 2 {
		count[cur - 'A'] += cnt - 2
	}
	fmt.Println(count)
	return count[0] > count[1]
}

func winnerOfGame(colors string) bool {
	n := len(colors)
	sumA,sumB := 0,0

	for i:=0;i<n;{
		countA,countB := 0,0
		for i < n && colors[i]=='A' {
			countA++
			i++
		}
		if countA > 2 {
			sumA += countA - 2
		}

		for i < n && colors[i]=='B' {
			countB++
			i++
		}
		if countB > 2 {
			sumB += countB - 2
		}
	}
	return sumA > sumB
}