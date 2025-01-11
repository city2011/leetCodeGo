package main

import "fmt"

func main() {
	res := groupAnagrams([]string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"},
	)
	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {
	count := map[byte]float64{
		'a': 2, 'b': 3, 'c': 5, 'd': 7, 'e': 11, 'f': 13, 'g': 17, 'h': 19, 'i': 23, 'j': 29, 'k': 31, 'l': 37, 'm': 41, 'n': 43,
		'o': 47, 'p': 53, 'q': 59, 'r': 61, 's': 67, 't': 71, 'u': 73, 'v': 79, 'w': 83, 'x': 89, 'y': 97, 'z': 101,
	}
	mem := map[float64][]string{}

	for _, str := range strs {
		var cnt float64 = 1
		for i, _ := range str {
			cnt *= count[str[i]]
		}
		fmt.Println(cnt)
		mem[cnt] = append(mem[cnt], str)
	}

	var res [][]string
	for _, strings := range mem {
		res = append(res, strings)
	}
	return res
}
