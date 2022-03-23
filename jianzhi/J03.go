package main

func main() {

}

func findRepeatNumber(nums []int) int {
	n := len(nums)
	visit := make([]bool, n)
	for _,num := range nums {
		if visit[num] {
			return num
		} else {
			visit[num] = true
		}
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			} else {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			}
		}
	}
	return -1
}