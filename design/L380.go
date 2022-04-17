package main

import (
	"math/rand"
)

func main() {

}

type RandomizedSet struct {
	mem map[int]int
	list []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		make(map[int]int),
		[]int{}}
}


func (this *RandomizedSet) Insert(val int) bool {
	if _,ok := this.mem[val]; ok {
		return false
	}
	this.list = append(this.list, val)
	this.mem[val] = len(this.list) - 1
	return true
}


func (this *RandomizedSet) Remove(val int) bool {
	id,ok := this.mem[val]
	if !ok {
		return false
	}
	last := len(this.list)
	this.list[id] = this.list[last - 1]
	this.mem[this.list[id]] = id
	delete(this.mem, val)
	this.list = this.list[:last - 1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	pos := rand.Intn(len(this.list))
	return this.list[pos]
}