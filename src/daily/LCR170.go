/*
在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的「交易逆序对」总数。

示例 1：
输入：record = [9, 7, 5, 4, 6]
输出：8
解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。

提示：
0 <= record.length <= 50000
*/
package main

type pair struct {
	val   int
	index int
}

func reversePairs(record []int) int {
	n := len(record)
	if n <= 1 {
		return 0
	}
	pairs := make([]pair, n)
	for i, r := range record {
		pairs[i] = pair{r, i}
	}

	_, count := mergeSort(pairs)
	return count
}

func mergeSort(paris []pair) ([]pair, int) {
	n := len(paris)
	if n == 1 {
		return paris, 0
	}
	mid := n / 2
	left, lcount := mergeSort(paris[:mid])
	right, rcount := mergeSort(paris[mid:])

	merged, mcount := merge(left, right)
	return merged, lcount + rcount + mcount
}

func merge(left, right []pair) ([]pair, int) {
	n, m := len(left), len(right)
	count := 0
	newP := make([]pair, 0)
	i, j := 0, 0
	for i < n && j < m {
		if left[i].val > right[j].val {
			newP = append(newP, right[j])
			count += n - i
			j++
		} else {
			newP = append(newP, left[i])
			i++
		}
	}
	if i < n {
		newP = append(newP, left[i:]...)
	}
	if j < m {
		newP = append(newP, right[j:]...)
	}
	return newP, count
}
