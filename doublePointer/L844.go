package main

import "fmt"

func main() {
	ans := backspaceCompare("a##c", "#a#c")
	fmt.Print(ans)
	fmt.Println("hello, go")
}

func backspaceCompare(S string, T string) bool {
	skipS, skipT := 0, 0
	i, j := len(S) - 1, len(T) - 1
	for i >= 0 {
		for i >= 0 {
			if S[i] == '#' {
				skipS++
				i--
			} else if skipS > 0{
				skipS--
				i--
			} else {
				break
			}
		}

		for j >= 0 {
			if T[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if S[i] != T[j] {
				return false
			}
		} else {
			if i >= 0 || j >= 0 {
				return false
			}
		}
		i--
		j--
	}
	return false
}
