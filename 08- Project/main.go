package main

import (
	"example.com/calc/job"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.01, 0.07, 0.15}
	for _, taxRate := range taxRates {
		inputFilePath := "job/prices.txt"
		outputFilePath := fmt.Sprintf("job/prices_with_tax_%.2f.json", taxRate)
		j := job.New(taxRate, inputFilePath, outputFilePath)
		j.Process()
	}
}