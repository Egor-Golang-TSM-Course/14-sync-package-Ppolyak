package main

import (
	"fmt"
	"sync"
)

//const workersNumber = 8

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (b *BankAccount) Deposit(amount int) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance = b.balance + amount
	//fmt.Println(b.balance)
	return b.balance
}

func (b *BankAccount) Withdraw(amount int) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	if amount > b.balance {
		fmt.Println("Withdraw bigger then your balance")
	} else {
		b.balance = b.balance - amount
	}
	return b.balance
}

func main() {
	b := BankAccount{
		balance: 20000,
	}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			b.Deposit(100)
			wg.Done()
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			b.Withdraw(100100)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(b.balance)
}
