package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.01, 0.07, 0.15}

	result := map[float64][]float64{}

	for _, taxRate := range taxRates {
		pricesWithTaxIncluded := []float64{}
		for _, price := range prices {
			pricesWithTaxIncluded = append(pricesWithTaxIncluded, (1 + taxRate) * (price))
		}
		result[taxRate] = pricesWithTaxIncluded
	}

	for key, values := range result {
		fmt.Printf("Tax rate : %.2f has the following prices :  ", key)
		for _, price := range values {
			fmt.Printf("%.2f,	", price)
		}
		fmt.Println()
	}
}