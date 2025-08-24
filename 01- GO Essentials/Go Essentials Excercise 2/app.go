package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)
const fileName = "balance.txt"

func main(){
	balance , err := readFromFile()
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

func printOptions(){
		fmt.Println("WELCOME TO BANK")
		fmt.Println("CHOOSE WHAT OPERATION YOU WANNA DO:")
		fmt.Println("1. BALANCE ENQUIRY")
		fmt.Println("2. WITHDRAWAL")
		fmt.Println("3. DEPOSIT")
		fmt.Println("4. EXIT")
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
	err := writeToFile(balance)
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
	err := writeToFile(balance)
	if err != nil {
		fmt.Println(err)
	}
	return balance
}


func writeToFile(balance float64) error {
	// convert balance to string
	stringBalance := fmt.Sprint(balance)
	// store balance in file, but at first convert string to list of bytes, 0644 is the permission
	err := os.WriteFile(fileName, []byte(stringBalance), 0644)
	if err != nil {
		return errors.New("FAILED TO WRITE BALANCE TO FILE")
	}
	return nil
}

func readFromFile() (float64, error){
	//read balance from file as bytes
	data, err :=os.ReadFile(fileName)
	// if an error occurred
	if err != nil {
		return 0, errors.New("FAILED TO READ BALANCE FILE")
	}
	// convert data to string
	stringBalance := string(data)
	// convert string to float64
	balance, err := strconv.ParseFloat(stringBalance, 64)
	if err != nil {
		return 0, errors.New("FAILED TO PARSE BALANCE")
	}
	return balance, nil

}