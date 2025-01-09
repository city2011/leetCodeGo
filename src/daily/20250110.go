package main

import "fmt"

func main() {
	res := validSubstringCount("dddddededddeeeddd", "eee")
	fmt.Println(res)
}

func validSubstringCount(word1 string, word2 string) int64 {
	diff := make([]int, 26)
	cnt := 0
	for i, _ := range word2 {
		diff[word2[i] - 'a']++
	}
	for _, d := range diff {
		if d > 0 {
			cnt++
		}
	}

	res,r := 0, 0
	for l := 0; l < len(word1); {
		for l := 0; l < len(word1); {
			for r < len(word1) && cnt > 0 {
				update(diff, &cnt, -1, word1[r])
				r++
			}
			if cnt == 0 {
				res += len(word1) - r + 1
			}
			update(diff, &cnt, 1, word1[l])
			l++
		}
		update(diff, &cnt, 1, word1[l])
		l++
	}
	return int64(res)
}

func update(diff []int, cnt *int, add int, w byte){
	diff[w - 'a'] += add
	if add == -1 && diff[w - 'a'] == 0 {
		*cnt--
	} else if add == 1 && diff[w - 'a'] == 1 {
		*cnt++
	}
}