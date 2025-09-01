package main

import (
	"fmt"
	"example.com/calc/filemanager"
	"example.com/calc/job"
	// "example.com/calc/cmdmanager"
)

func main() {
	taxRates := []float64{0, 0.01, 0.07, 0.15}
	for _, taxRate := range taxRates {
		inputFilePath := "job/prices.txt"
		outputFilePath := fmt.Sprintf("job/prices_with_tax_%.2f.json", taxRate)
		fm := filemanager.New(inputFilePath, outputFilePath)
		// cmdM := cmdmanager.New()
		j := job.New(taxRate, fm)
		err := j.Process()
		if err != nil {
			fmt.Println("Error processing job:", err)
			return 
		}
	}
}
