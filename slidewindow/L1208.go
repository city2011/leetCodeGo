package main

func main() {

}

func equalSubstring(s string, t string, maxCost int) int {
	l,diff,r := 0,0,0
	for ;r < len(s);r++ {
		diff += absDiff(s[r], t[r])
		if diff > maxCost {
			diff -= absDiff(s[l], t[l])
			l++
		}
	}
	return r - l + 1
}

func absDiff(a,b byte) int {
	if a > b {
		return int(a - b)
	} else {
		return int(b - a)
	}
}

func max(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}