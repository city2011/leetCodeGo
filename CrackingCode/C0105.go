package main

import "fmt"

func main() {
	fmt.Println(oneEditAway("abd", "abc"))
	fmt.Println(oneEditAway("abde", "abc"))
	fmt.Println(oneEditAway("abde", "abd"))
	fmt.Println(oneEditAway("abd", "abdf"))

}

func oneEditAway(first string, second string) bool {
	if len(first) == len(second) {
		return disZeroEditable(first, second)
	} else if len(first) - len(second) == 1 {
		return disOneEditable(first, second)
	} else if len(second) - len(first) == 1 {
		return disOneEditable(second, first)
	} else {
		return false
	}
}

func disOneEditable(first string, second string) bool{
	var sum int
	n := len(first)
	for i,j := 0,0; i<n && j<n-1; {
		if second[j] != first[i] {
			if sum >= 1 {
				return false
			}
			sum++
			i++
		} else {
			i++
			j++
		}
	}
	return sum <= 1
}

func disZeroEditable(first string, second string) bool{
	var sum int
	for i := 0;i<len(first);i++ {
		if first[i] != second[i] {
			sum++
		}
	}
	return sum <= 1
}