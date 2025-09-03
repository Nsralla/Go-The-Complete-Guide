package main

import (
	"fmt"
	"example.com/calc/filemanager"
	"example.com/calc/job"
	// "example.com/calc/cmdmanager"
)

func main() {
	taxRates := []float64{0, 0.01, 0.07, 0.15}
	doneChan := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {
		doneChan[index] = make(chan bool)
		errorChans[index] = make(chan error)
		inputFilePath := "job/prices.txt"
		outputFilePath := fmt.Sprintf("job/prices_with_tax_%.2f.json", taxRate)
		fm := filemanager.New(inputFilePath, outputFilePath)
		// cmdM := cmdmanager.New()
		j := job.New(taxRate, fm)
		go j.Process(doneChan[index], errorChans[index])
		// if err != nil {
		// 	fmt.Println("Error processing job:", err)
		// 	return 
		// }
	}
 
	for index := range taxRates {
		select { //  wait for error, or success case.
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println("Error processing job:", err)
				// Stop the program if there's an error
				return
			}
		case <-doneChan[index]:
			// Job completed successfully
			fmt.Printf("Job %d completed successfully\n", index)
		}
	}

	// // don't exit until all jobs write on the channel
	// for _, channel := range doneChan {
	// 	<-channel   
	// }
}
