package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type InputItem struct {
	ID   int
	Data int
}

type ProcessedOutput struct {
	ID     int
	Result int
}

func main() {
	// Create seed
	rand.Seed(time.Now().UnixNano())

	// Create wg
	var wg sync.WaitGroup

	totalWorkers := 3
	totalItemsToProcess := 9

	jobsChan := make(chan InputItem)
	resultsChan := make(chan ProcessedOutput)

	// Read data from terminal, then send to jobsChannel
	go func(){
		id := 0
		data := 0
		for range totalItemsToProcess{
			fmt.Println("\nEnter job Id and Job data : ")
			fmt.Scanln(&id)
			fmt.Scanln(&data)
			job := InputItem{ID: id, Data: data}
			jobsChan <- job
		}
		close(jobsChan)
	}()

	for i:=0; i<totalWorkers; i++{
		wg.Add(1)
		go worker(&wg,resultsChan, jobsChan)
	}

	for i:=0; i<totalItemsToProcess;i++{
		result := <- resultsChan
		fmt.Printf("Result %d has been received with result = %d\n",result.ID, result.Result)
	}


	wg.Wait()
	close(resultsChan)

}

func worker(wg *sync.WaitGroup, resultsChan chan<- ProcessedOutput, jobsChan <- chan InputItem) {
	defer wg.Done()
	// Read input items (channel)
	min := 100
	max := 700
	for job := range jobsChan{
		randomNumber := rand.Intn(max - min + 1) + min
		time.Sleep(time.Millisecond * time.Duration(randomNumber))
		output := ProcessedOutput{
			ID: job.ID,
			Result: job.Data + 100,
		}
		resultsChan <- output
	}
	// small delay
	//perform simple  transformation
	// send output (channel)
}
