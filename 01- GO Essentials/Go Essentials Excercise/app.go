package main

import (
	"errors"
	"fmt"
	"os"
)

func main(){

	revenue, expenses, taxRate, err := readVariables()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	EBT, profit, ratio := CalcValues(revenue, expenses, taxRate)
	stringEBT := fmt.Sprintf("%f", EBT) // to convert float64 to string
	stringProfit := fmt.Sprintf("%f", profit)
	stringRatio := fmt.Sprintf("%f", ratio)
    fullText := stringEBT + "\n" + stringProfit + "\n" + stringRatio
	os.WriteFile("values.txt", []byte(fullText), 0644)

	// fmt.Print("Earnings before tax: ")
	// fmt.Println(EBT)
	// fmt.Print("Earnings after tax: ")
	// fmt.Println(profit)
	// fmt.Printf("Ratio: %.2f", ratio)

}

func readVariables()(revenue float64, expenses float64, taxRate float64, error error){

	fmt.Println("Profit Calculator")
	fmt.Print("Enter the total revenue: ")
	fmt.Scan(&revenue)
	//validate revenue
	if revenue < 0 {
		fmt.Println("Invalid revenue. Please enter a positive value.")
		return 0, 0, 0, errors.New("invalid revenue: revenue must be positive")
	}
	fmt.Print("Enter the total expenses: ")
	fmt.Scan(&expenses)
	//validate expenses
	if expenses < 0 {
		fmt.Println("Invalid expenses. Please enter a positive value.")
		return 0, 0, 0, errors.New("invalid expenses: expenses must be positive")
	}
	fmt.Print("Enter the tax rate (in %): ")
	fmt.Scan(&taxRate)
	// validate tax rate
	if taxRate < 0 || taxRate > 100 {
		fmt.Println("Invalid tax rate. Please enter a value between 0 and 100.")
		return 0, 0, 0, errors.New("invalid tax rate: tax rate must be between 0 and 100")
	}
	return revenue, expenses, taxRate, nil
}

func CalcValues(revenue, expenses, taxRate float64)(EBT float64, profit float64, ratio float64){
		EBT = revenue - expenses
		profit = EBT * (1 - taxRate/100)
		ratio = EBT / profit
		return
}