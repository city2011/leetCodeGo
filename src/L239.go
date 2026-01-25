package src

// 1,2,3,4,5,6,7 3
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	res := make([]int, len(nums)-k+1)
	que := make([]int, 0)

	for j, i := 0, 1-k; j < len(nums); j, i = j+1, i+1 {
		if i > 0 && nums[i-1] == que[0] {
			que = que[1:]
		}

		for len(que) > 0 && nums[j] > que[len(que)-1] {
			que = que[:len(que)-1]
		}

		que = append(que, nums[j])

		if i >= 0 {
			res[i] = que[0]
		}
	}

	return res
}

// 1,2,3,4,5,6,7 3
func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	que := make([]int, 0)
	res := make([]int, len(nums)-k+1)

	for i, j := 1-k, 0; j < len(nums); i, j = i+1, j+1 {
		// 出
		if i > 0 && i-1 == que[0] {
			que = que[1:]
		}

		// 入 单调队列，当前元素入，比它小的出
		for len(que) > 0 && nums[j] >= nums[que[len(que)-1]] {
			que = que[:len(que)-1]
		}
		que = append(que, j)

		if i >= 0 {
			res[i] = nums[que[0]]
		}
	}

	return res
}
