package main

import "fmt"

func main() {
	res := lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext")
	fmt.Println(res)
	fmt.Println(len("\n\t\n\t"))
}

func lengthLongestPath(input string) int {
	isFile := false
	leng := 0
	level := 0
	oldlevel := 0
	ans := 0
	var rec []int
	for i:=0;i<len(input); i++{
		if input[i] != '\n'{
			leng++
			if input[i] == '.' {
				isFile = true
			}
		} else {
			// qing suan
			// rec = append(rec, leng)
			if oldlevel >= level {
				rec = rec[:level]
			}
			rec = append(rec, leng)
			if isFile {
				recall := sum(rec)
				ans = max(ans, recall + level)
			}

			oldlevel = level
			k := i + 1
			level = 0
			for k < len(input) && input[k]=='\t' {
				level += 1
				k++
			}
			// calc level
			i = k - 1
			leng = 0
			isFile = false
		}
	}

	if leng > 0 && isFile {
		if oldlevel >= level {
			rec = rec[:level]
		}
		rec = append(rec, leng)
		recall := sum(rec)
		ans = max(ans, recall + level)
	}
	return ans
}

func max(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 可以用前缀和来改进sum
func sum(arr []int) int {
	sum := 0
	for _,a := range arr {
		sum += a
	}
	return sum
}