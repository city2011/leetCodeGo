package main

func main() {

}

func isAlienSorted(words []string, order string) bool {
	index := make(map[byte]int, 26)
	for i := 0; i < 26; i++ {
		index[order[i]] = i
	}
	for i := 0; i < len(words) - 1; i++ {
		if !compare(words[i], words[i+1], index) {
			return false
		}
	}
	return true
}

func compare(a string, b string, idx map[byte]int) bool{
	for i := 0; i < len(a); i++ {
		if i >= len(b) {
			return false
		}
		if idx[a[i]] < idx[b[i]] {
			return true
		}
		if idx[a[i]] > idx[b[i]] {
			return false
		}

 	}
 	return true
}
