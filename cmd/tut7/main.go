package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int = 1000

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: Withdraw function
	go func() {
		defer wg.Done()
		withdrawMoney(500)
	}()

	// Goroutine 2: Deposit function
	go func() {
		defer wg.Done()
		depositMoney(300)
	}()

	wg.Wait()
	fmt.Printf("Final balance: %d\n", balance)
}

func withdrawMoney(amount int) {
	// Simulate reading and checking balance
	currentBalance := balance

	// Simulate some processing time to increase race condition likelihood
	time.Sleep(10 * time.Millisecond)

	// Potentially incorrect withdrawal
	if currentBalance >= amount {
		balance = currentBalance - amount
	}
}

func depositMoney(amount int) {
	// Simulate reading current balance
	currentBalance := balance

	// Simulate processing time
	time.Sleep(5 * time.Millisecond)

	// Add deposit
	balance = currentBalance + amount
}
