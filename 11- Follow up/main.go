package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	balance := 1000
	var wg sync.WaitGroup
	mu := sync.Mutex{}

		// Create a new random source
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	wg.Add(10)
	for i := 0; i < 5; i++ {
	
		// Generate random number between 500 and 1500
		randomNum := r.Intn(1001) + 500  // 500 to 1500 inclusive
		go Deposit(randomNum, &balance, &mu, &wg)
		go Withdrawl(randomNum, &balance, &mu, &wg)
	}

	wg.Wait()
	fmt.Println("Balance: ", balance)

}

func Deposit(amount int, balance *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	
	
	mu.Lock()
	defer wg.Done()
	(*balance) += amount
	fmt.Printf("Depositing %v. your balance is %v \n", amount, *balance)
	mu.Unlock()

}

func Withdrawl(amount int, balance *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	defer wg.Done()
	if (*balance) - amount < 0 {
		fmt.Printf("You don't have sufficent amount to withdraw %v, your balance is %v\n", amount, *balance)
	} else {
		
		(*balance) -= amount
		fmt.Printf("Withdrawling %v, your balance is %v\n", amount, *balance)
	}
	mu.Unlock()
}
