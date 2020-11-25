package main

import "fmt"

func main() {
	sortString("dfadfasf")
}

func sortString(s string) string {
	cnt := [26]int{}
	for _, c := range s {
		fmt.Println(c - 'a')
		cnt[c-'a']++
	}
	n := len(s)
	ans := make([]byte, 0, n)
	fmt.Println(len(ans))
	for len(ans) < n {
		for i := byte('a'); i <= 'z'; i++ {
			if (cnt[i-'a'] > 0) {
				ans = append(ans, i)
				cnt[i-'a']--
			}
		}
		for i := byte('z'); i >= 'a'; i-- {
			if (cnt[i-'a'] > 0) {
				ans = append(ans, i)
				cnt[i-'a']--
			}
		}
	}

	return string(ans)
}

func sortStringLeetCode(s string) string {
	cnt := ['z' + 1]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	n := len(s)
	ans := make([]byte, 0, n)
	for len(ans) < n {
		for i := byte('a'); i <= 'z'; i++ {
			if cnt[i] > 0 {
				ans = append(ans, i)
				cnt[i]--
			}
		}
		for i := byte('z'); i >= 'a'; i-- {
			if cnt[i] > 0 {
				ans = append(ans, i)
				cnt[i]--
			}
		}
	}
	return string(ans)
}
