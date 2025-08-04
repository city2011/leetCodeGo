package main

import "fmt"

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	m1    map[int]int
	m2    map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for _, v := range nums1 {
		m1[v]++
	}
	for _, v := range nums2 {
		m2[v]++
	}
	return FindSumPairs{
		nums1: nums1,
		nums2: nums2,
		m1:    m1,
		m2:    m2,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	this.m2[this.nums2[index]]--
	this.nums2[index] = this.nums2[index] + val
	this.m2[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	sum := 0
	for k, v := range this.m1 {
		if value, ok := this.m2[tot-k]; ok {
			sum += v * value
		}
	}
	return sum
}

func main() {
	obj := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	fmt.Println(obj.Count(7))
	obj.Add(3, 2)
	fmt.Println(obj.Count(8))
	fmt.Println(obj.Count(4))
	obj.Add(1, 1)
	obj.Add(0, 1)
	fmt.Println(obj.Count(7))
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
