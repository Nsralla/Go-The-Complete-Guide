package main

import (
	"fmt"
	"example.com/bank/utilities"
)
const fileName = "balance.txt"

func main(){
	balance , err := utilities.ReadFromFile(fileName)
	if err != nil {
		fmt.Println(err)
		// return
		panic(err)
	}

	var choice int
	for{
	    printOptions()
		// read user input
		fmt.Scan(&choice)
		// check user input
		if choice == 1 {
			fmt.Println("YOUR BALANCE IS: ", balance)
		}else if choice == 2 {
			balance = withdraw(balance)
			fmt.Println("YOUR BALANCE IS: ", balance)
		}else if choice == 3 {
			balance = deposit(balance)
			fmt.Println("YOUR BALANCE IS: ", balance)
		}else if choice == 4 {
			fmt.Println("THANK YOU FOR BANKING WITH US")
			return
		}else{
			fmt.Println("INVALID CHOICE")
			continue
		}

	}

}



func withdraw(balance float64) float64 {
	var amount float64
	fmt.Print("Enter amount to withdraw: ")
	fmt.Scan(&amount)
	if amount > balance {
		fmt.Println("INSUFFICIENT BALANCE")
		return balance
	}else if amount <=0 {
		fmt.Println("INVALID AMOUNT")
		return balance
	}
	balance -= amount
	err := utilities.WriteToFile(balance, fileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return balance
}

func deposit(balance float64) float64 {
	var amount float64
	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("INVALID AMOUNT")
		return balance
	}
	balance += amount
	err := utilities.WriteToFile(balance, fileName)
	if err != nil {
		fmt.Println(err)
	}
	return balance
}


