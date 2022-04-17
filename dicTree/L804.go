package main

import "fmt"

func main() {
	res := uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})
	fmt.Println(res)
}

var sum int
var mosCode = []string {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}


type dicTree struct {
	children [2]*dicTree
	end bool
}

func trans2mosCode(str string) (ans []int){
	for i:=0;i<len(str);i++ {
		for _,mosStr := range mosCode[str[i] - 'a'] {
			if mosStr == '.' {
				ans = append(ans, 1)
			} else {
				ans = append(ans, 0)
			}
		}
	}
	return ans
}

func uniqueMorseRepresentations(words []string) int {
	var mosint = make([][]int, len(words))
	idx := 0
	for _,str := range words {
		mosint[idx] = trans2mosCode(str)
		idx++
	}

	var dic dicTree

	for _,each := range mosint {
		dic.add(each)
	}

	return sum
}

func (t *dicTree) add(mosint []int) {
	for _,num := range mosint {
		if t.children[num] == nil {
			t.children[num] = &dicTree{}
		}
		t = t.children[num]
	}
	if t.end != true {
		t.end = true
		sum++
	}
}

