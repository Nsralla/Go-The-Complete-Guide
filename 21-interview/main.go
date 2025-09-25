package main 

import (
	"fmt"
	"sync"
)


type NumSquarer struct {
	numWorkers int
	// Add other fields as needed
}

// NewNumSquarer initializes the NumSquarer with a given number of workers.
func NewNumSquarer(numWorkers int) *NumSquarer {
	return &NumSquarer{
		numWorkers: numWorkers,
	}
}

// Process spawns worker goroutines and squares all numbers in 'numbers'.
// Returns the results slice with squares in some defined order.
func (ns *NumSquarer) Process( wg *sync.WaitGroup, resultsChan chan int, inputChan <-chan int) []int {
	// TODO: Implement logic:

	for i:=0;i<ns.numWorkers;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			for number := range inputChan{
				// write to output channel
				resultsChan <- number * number
			}
		}()
	}

	wg.Wait()
	close(resultsChan)
	// create slice to hold results
	results := []int{}
	for result := range resultsChan{
		results = append(results, result)
	}
	return results
}

// main demonstrates how to use NumSquarer.
func main() {
	var wg sync.WaitGroup
	ns := NewNumSquarer(3)           // 3 workers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Input channel
	inputChan := make(chan int, len(numbers))
	for _, num := range numbers{
		inputChan <- num
	}
	close(inputChan) // close the input channel.
	// Output channel
	resultsChan := make(chan int, len(numbers))
	results := ns.Process( &wg, resultsChan, inputChan)
	
	
	fmt.Println("Input:", numbers)
	fmt.Println("Squares:", results)




}