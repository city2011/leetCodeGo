package main

import "unicode"

func main() {

}

func mostCommonWord(paragraph string, banned []string) string {
	banmap := make(map[string]bool)
	for _,ban := range banned {
		banmap[ban] = true
	}

	wordm := make(map[string]int)

	var word []byte
	word = nil
	for i,n := 0, len(paragraph); i <= n; i++ {
		if i < n && unicode.IsLetter(rune(paragraph[i])) {
			word = append(word, byte(unicode.ToLower(rune(paragraph[i]))))
		} else if word != nil {
			str := string(word)
			if ok := banmap[str]; !ok {
				wordm[str]++
			}
			word = nil
		}
	}

	max := -1
	var res string
	for k,v := range wordm {
		if v > max {
			max = v
			res = k
		}
	}
	return res
}