package main

func smallestSubsequence(s string) string {
	left := ['z' + 1]int{}
	already := ['z' + 1]bool{}

	var ans []rune
	for _, c := range s {
		left[c]++
	}

	for _, c := range s {
		left[c]--
		if already[c] {
			continue
		}

		for len(ans) > 0 && c < ans[len(ans)-1] && left[ans[len(ans)-1]] > 0 {
			already[ans[len(ans)-1]] = false
			ans = ans[:len(ans)-1]
		}
		ans = append(ans, c)
		already[c] = true
	}

	return string(ans)
}
