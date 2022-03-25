package main

import "fmt"

func main() {
	str := pushDominoes(".L.R...LR..L..")
	fmt.Println(str)
}

func pushDominoes(dominoes string) string {
	if len(dominoes) == 0 {
		return ""
	}
	cur := 0
	left := -1
	var res string
	for len(dominoes) > cur {
		if dominoes[cur] == 'L' {
			left = cur
		} else if dominoes[cur] == 'R' {
			break
		}
		cur++
	}

	for i := 0; i < left+1; i++ {
		res += "L"
	}

	for i := left + 1; i < cur && i < len(dominoes); i++ {
		res += "."
	}

	if cur >= len(dominoes) {
		return res
	}

	mid := cur
	right := cur
	cur++
	for len(dominoes) > cur {
		if dominoes[cur] == 'R' {
			right = cur
		} else if dominoes[cur] == 'L' {
			break
		}
		cur++
	}

	if cur >= len(dominoes) {
		for i := mid; i < len(dominoes); i++ {
			res += "R"
		}
		return res
	}

	newLeft := right + (cur-right-1)/2
	newRight := cur - (cur-right-1)/2

	for i := mid; i <= newLeft; i++ {
		res += "R"
	}
	if newRight-newLeft == 2 {
		res += "."
	}
	for i := newRight; i <= cur; i++ {
		res += "L"
	}

	return res + pushDominoes(dominoes[cur+1:])
}
