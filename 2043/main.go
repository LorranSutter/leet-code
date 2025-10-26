package main

import "fmt"

type Bank struct {
	NumAccounts int
	Balances    []int64
}

func Constructor(balance []int64) Bank {
	balances := make([]int64, len(balance)+1)

	for i, b := range balance {
		balances[i+1] = b
	}

	return Bank{NumAccounts: len(balances), Balances: balances}
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 < 1 || account2 < 1 || account1 > b.NumAccounts || account2 > b.NumAccounts {
		return false
	}
	if b.Balances[account1] < money {
		return false
	}

	b.Balances[account1] -= money
	b.Balances[account2] += money
	return true
}

func (b *Bank) Deposit(account int, money int64) bool {
	if account < 1 || account > b.NumAccounts {
		return false
	}

	b.Balances[account] += money
	return true
}

func (b *Bank) Withdraw(account int, money int64) bool {
	if account < 1 || account > b.NumAccounts {
		return false
	}
	if b.Balances[account] < money {
		return false
	}

	b.Balances[account] -= money
	return true
}

func main() {
	obj := Constructor([]int64{10, 100, 20, 50, 30, 0})
	fmt.Println(obj.Withdraw(3, 10) == true)
	fmt.Println(obj.Transfer(5, 1, 20) == true)
	fmt.Println(obj.Deposit(5, 20) == true)
	fmt.Println(obj.Transfer(3, 4, 15) == false)
	fmt.Println(obj.Withdraw(10, 50) == false)
	fmt.Println(obj.Deposit(6, 20) == true)
	fmt.Println(obj.Deposit(99, 20) == false)
}
