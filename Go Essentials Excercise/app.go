package main
import (
	"fmt"
)

func main(){
	revenue, expenses, taxRate := readVariables()
    EBT, profit, ratio := CalcValues(revenue, expenses, taxRate)
	fmt.Print("Earnings before tax: ")
	fmt.Println(EBT)
	fmt.Print("Earnings after tax: ")
	fmt.Println(profit)
	fmt.Printf("Ratio: %.2f", ratio)

}

func readVariables()(revenue float64, expenses float64, taxRate float64){

	fmt.Println("Profit Calculator")
	fmt.Print("Enter the total revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the total expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate (in %): ")
	fmt.Scan(&taxRate)
	return revenue, expenses, taxRate
}

func CalcValues(revenue, expenses, taxRate float64)(EBT float64, profit float64, ratio float64){
		EBT = revenue - expenses
		profit = EBT * (1 - taxRate/100)
		ratio = EBT / profit
		return
}