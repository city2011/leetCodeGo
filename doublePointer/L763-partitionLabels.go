package main

import "fmt"

func main() {
	par := partitionLabels("abdkldddfgadgiojhwenw")
	fmt.Print(par)
}

func partitionLabels(S string) (partition []int) {
	lastPos := [26]int{}

	for i := 0; i < len(S); i++ {
		pos := S[i] - 'a'
		lastPos[pos] = i
	}

	start, end := 0, 0
	for i, c := range S {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}
