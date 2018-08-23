package main

import (
	"fmt"
	"log"
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

// Withdraw 取款
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()

	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

// Deposit 存款
func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

// 余额
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func deposit(amount int) {
	balance += amount
}

func main() {
	var n sync.WaitGroup

	for i := 0; i < 10; i++ {
		n.Add(1)
		go func(i int) {
			defer n.Done()
			Deposit(i)
		}(i)
	}

	for j := 0; j < 20; j++ {
		n.Add(1)
		go func(j int) {
			defer n.Done()
			if b := Withdraw(j); !b {
				log.Printf("Withdraw failed! amount = %d\n", j)
			}
		}(j)
	}

	n.Wait()
	fmt.Println(balance)
}
