package main

func main() {

}

func shortestToChar(s string, c byte) (ans []int) {
	if len(s) == 0 {
		return ans
	}
	var rec []int
	ans = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			rec = append(rec, i)
		}
	}
	if len(rec) == 0 {
		return ans
	}

	for i,k := rec[0],0; i >= 0; i,k = i-1, k+1 {
		ans[i] = k
	}

	for i,k := rec[len(rec) - 1],0; i < len(ans); i,k = i+1, k+1 {
		ans[i] = k
	}

	for i := 1; i < len(rec) - 1; i++ {
		l := rec[i]
		r := rec[i + 1]
		k := 0
		for l <= r {
			ans[l],ans[r] = k,k
			k++
			l++
			r--
		}
	}
	return ans
}
