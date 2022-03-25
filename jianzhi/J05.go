package main

import "fmt"

func main() {
	fmt.Println(replaceSpace2("We are happy."))
}

func replaceSpace(s string) string {
	var res string
	for _, byte1 := range s {
		if byte1 == ' ' {
			res += "%20"
		} else {
			res += string(byte1)
		}
	}
	return res
}

func replaceSpace2(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	// 计算空格数量
	for _, v := range b {
		if v == ' ' {
			spaceCount++
		}
	}
	// 扩展原有切片
	resizeCount := spaceCount * 2
	tmp := make([]byte, resizeCount)
	b = append(b, tmp...)
	i := length - 1
	j := len(b) - 1
	for i < j {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j - 3
		}
	}
	return string(b)
}
