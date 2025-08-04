package main

import "fmt"

func totalFruit(fruits []int) int {
	cnt := map[int]int{}
	res := 0
	left, right := 0, 0
	for _, v := range fruits {
		right++
		cnt[v]++
		for len(cnt) > 2 {
			out := fruits[left]
			cnt[out]--
			if cnt[out] == 0 {
				delete(cnt, out)
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func main() {
	res := totalFruit([]int{1, 2, 3, 2, 2})
	fmt.Println(res)
}
