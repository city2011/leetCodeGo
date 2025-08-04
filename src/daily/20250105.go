package main

import (
	"fmt"
	"leetCodeGo/src/basic"
)

var money = [5]int{20, 50, 100, 200, 500}
var bad = []int{-1}

type ATM struct {
	sto [5]int
}

func (a *ATM) Moneys() [5]int {
	return a.sto
}

func Constructor() ATM {
	return ATM{
		sto: [5]int{},
	}
}

func (a *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < len(a.sto); i++ {
		a.sto[i] = a.sto[i] + banknotesCount[i]
	}
}

func (a *ATM) Withdraw(amount int) []int {
	withdraw := []int{0, 0, 0, 0, 0}
	res := amount
	for i := len(a.sto) - 1; i >= 0; i-- {
		if res == 0 {
			break
		}
		if a.sto[i] != 0 && res >= money[i] {
			out := int(basic.Min(res/money[i], a.sto[i]))
			withdraw[i] = out
			res -= out * money[i]
		}
	}
	if res != 0 {
		return bad
	} else {
		for i := 0; i < len(a.sto); i++ {
			a.sto[i] -= withdraw[i]
		}
		return withdraw
	}
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */

func main() {
	obj := Constructor()
	obj.Deposit([]int{0, 0, 1, 2, 1})
	fmt.Println("money: ", obj.Moneys())
	fmt.Println(obj.Withdraw(600))
	fmt.Println("money: ", obj.Moneys())
	obj.Deposit([]int{0, 1, 0, 1, 1})
	fmt.Println("money: ", obj.Moneys())
	fmt.Println(obj.Withdraw(600))
	fmt.Println("money: ", obj.Moneys())
	fmt.Println(obj.Withdraw(550))
	fmt.Println("money: ", obj.Moneys())
}
