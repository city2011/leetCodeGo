package main

func main() {

}

func firstUniqChar(s string) byte {
	mem := make(map[int32]bool)
	for _,c := range s {
		if _,ok := mem[c];ok {
			mem[c] = true
		} else {
			mem[c] =false
		}
	}
	for _,c := range s {
		if !mem[c] {
			return byte(c)
		}
	}
	return byte(' ')
}