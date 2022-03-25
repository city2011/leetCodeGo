package main

func main() {

}

func reverseOnlyLetters(s string) string {
	ss := []byte(s)
	for i, j := 0, len(ss)-1; i < j; {
		for !isDigital(ss[i]) {
			i++
		}
		for !isDigital(ss[j]) {
			j--
		}
		if i <= j {
			tmp := ss[i]
			ss[i] = ss[j]
			ss[j] = tmp
			i++
			j--
		}
	}
	return string(ss)
}

func isDigital(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}
