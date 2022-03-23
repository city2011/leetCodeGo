package main

func main() {

}

func missingNumber(nums []int) int {
	var res int
	for _,num := range nums {
		res = res ^ num
	}
	for i := 0; i <= len(nums); i++ {
		res = res ^ i
	}
	return res
}

func missingNumber2(nums []int) int {
	l,r := 0, len(nums) - 1
	var m int

	for l <= r {
		m = l + (r - l) / 2
		if nums[m] > m {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return l
}