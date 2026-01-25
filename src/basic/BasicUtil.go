package basic

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func Gcd(a,b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}