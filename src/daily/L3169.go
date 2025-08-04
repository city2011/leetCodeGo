package main

import "sort"

/*
1 <= days <= 10^9
1 <= meetings.length <= 10^5
meetings[i].length == 2
1 <= meetings[i][0] <= meetings[i][1] <= days
*/

/*
解题思路：
 1. 排序，按照会议开始时间排序
 2. 计算会议开始时间之前的天数，累加到sumDays中
 3. 初始化right，用于记录目前‘合并会议’的最大结束时间。（‘合并会议’是指多个会议区间重叠，形成合并的会议区间）
 4. 遍历会议
    4.1 当前会议的开始时间比right + 1 还大，无法合并到当前的‘合并会议’。此时计算【空闲】天数，并累加到sumDays中
    4.2 当前会议的开始时间比right + 1 小，可以合并到当前的‘合并会议’。此时更新right的值为当前会议的结束时间
 4. 遍历结束，计算right后的空闲天数，并累加到sumDays中
 5. 此时空闲天数全部计算完毕，返回sumDays
*/
func countDays(days int, meetings [][]int) int {
	sumDays := 0
	// 题目有约束条件，这个其是可以不加。
	// 像后面的meeting本身的约束，和days本身的约束，都未做处理，如果是工程类的，一个是需要做基本的参数校验。
	if len(meetings) == 0 {
		return days
	}
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	if meetings[0][0] > 1 {
		sumDays += meetings[0][0] - 1
	}
	right := meetings[0][1]
	for _, meeting := range meetings {
		if meeting[0] > right+1 {
			sumDays += meeting[0] - right - 1
			right = meeting[1]
		} else {
			right = max(right, meeting[1])
		}
	}
	sumDays += days - right
	return sumDays
}

func main() {
	res := countDays(10, [][]int{{1, 3}, {2, 6}, {3, 4}, {3, 5}, {4, 6}})
	println(res)

	res = countDays(10, [][]int{})
	println(res)
}
