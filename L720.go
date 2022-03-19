package main

import (
	"fmt"
	"sort"
)

func main() {
	res2 := longestWord2([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"})
	res := longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"})
	fmt.Println(res, res2)
}

func longestWord(words []string) string {
	var res string
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j]) || len(words[i]) == len(words[j]) && words[i] > words[j]
	})

	mem := map[string]string {"": "ok"}

	for _, word := range words {
		if _,ok := mem[word[:len(word) - 1]];ok {
			res = word
			mem[word] = "ok"
		}
	}
	return res
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil || !node.children[ch].isEnd {
			return false
		}
		node = node.children[ch]
	}
	return true
}

func longestWord2(words []string) (longest string) {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	for _, word := range words {
		if t.Search(word) && (len(word) > len(longest) || len(word) == len(longest) && word < longest) {
			longest = word
		}
	}
	return longest
}
