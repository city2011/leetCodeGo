package _0250624

import "leetCodeGo/src/basic"

func findKDistantIndices(nums []int, key int, k int) []int {
	var ans []int
	end := len(nums) - 1
	l, r := 0, end
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			l = basic.Max(l, i-k)
			r = basic.Min(r, i+k)
			for j := l; j <= r; j++ {
				ans = append(ans, j)
			}
			l = r + 1
			r = end
		}
	}
	return ans
}
