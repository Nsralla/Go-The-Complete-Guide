package main

import (
	"fmt"
	"math"
)
const inflationRate = 2.5
func main() {
	
	var investmentAmount float64
	var returnRate float64
	var years float64


	fmt.Println("Investment Calculator")
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)
	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&returnRate)

	featureValue, futureRealValue := calcFutureValues(investmentAmount, returnRate, years)

	// fmt.Printf("The future value of the investment is: %.2f\n", featureValue)
	// Printf: print the text with special format
	// fmt.Printf("The future real value of the investment is: %.2f\n", futureRealValue)

	formattedFutureValue := fmt.Sprintf("THE FUTURE VALUE OF THE INVESTMENT IS : %.2f", featureValue)
	formattedFutureRealValue := fmt.Sprintf("THE FUTURE REAL VALUE OF THE INVESTMENT IS : %.2f", futureRealValue)
	fmt.Println(formattedFutureValue)
	fmt.Println(formattedFutureRealValue)
}

func calcFutureValues(investmentAmount, returnRate, years float64) (float64, float64){
	featureValue := investmentAmount * math.Pow((1+returnRate/100), years)
	// := this will create a new variable and assign the value to it
	// = this will assign the value to an existing variable
	futureRealValue := featureValue / math.Pow(1+inflationRate/100, years)
	return featureValue, futureRealValue
}

//method2:
// func calcFutureValues(investmentAmount, returnRate, years float64) (featureValue float64,featureRealValue float64){
// 	featureValue = investmentAmount * math.Pow((1+returnRate/100), years)
// 	futureRealValue = featureValue / math.Pow(1+inflationRate/100, years)
// 	return 
// }
