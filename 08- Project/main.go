package main

import (
	"example.com/calc/job"
)

func main() {
	taxRates := []float64{0, 0.01, 0.07, 0.15}

	// result := map[float64][]float64{}

	for _, taxRate := range taxRates {
		j := job.New(taxRate)
		j.Process()
	}

	
}