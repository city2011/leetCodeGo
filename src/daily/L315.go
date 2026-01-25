/*
给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

示例 1：
输入：nums = [5,2,6,1]
输出：[2,1,1,0]
解释：
5 的右侧有 2 个更小的元素 (2 和 1)
2 的右侧仅有 1 个更小的元素 (1)
6 的右侧有 1 个更小的元素 (1)
1 的右侧有 0 个更小的元素

示例 2：
输入：nums = [-1]
输出：[0]

示例 3：
输入：nums = [-1,-1]
输出：[0,0]
*/
package main

func countSmaller(nums []int) []int {
	type pair struct {
		idx int
		val int
	}
	pairs := make([]pair, len(nums))
	for i, num := range nums {
		pairs[i] = pair{i, num}
	}

	ans := make([]int, len(nums))

	var merge func(left, right []pair) []pair
	merge = func(left, right []pair) []pair {
		n, m := len(left), len(right)
		i, j := 0, 0
		newP := make([]pair, 0, n+m)
		for i < n && j < m {
			if left[i].val <= right[j].val {
				newP = append(newP, left[i])
				ans[left[i].idx] += j
				i++
			} else {
				newP = append(newP, right[j])
				j++
			}
		}
		if i < n {
			newP = append(newP, left[i:]...)
			for i < n {
				ans[left[i].idx] += j
				i++
			}
		}
		if j < m {
			newP = append(newP, right[j:]...)
		}
		return newP
	}

	var mergeSort func(pairs []pair) []pair
	mergeSort = func(pairs []pair) []pair {
		if len(pairs) == 1 {
			return pairs
		}
		mid := len(pairs) / 2
		left := mergeSort(pairs[:mid])
		right := mergeSort(pairs[mid:])

		return merge(left, right)
	}

	mergeSort(pairs)
	return ans
}
