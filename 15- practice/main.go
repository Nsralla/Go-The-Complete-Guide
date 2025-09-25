package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Item struct {
	ID   int
	Data string
}
type ProcessedItem struct {
	ID     int
	Result string
}

func worker(jobsChan <- chan Item, resultsChan chan <- ProcessedItem, wg * sync.WaitGroup, id int) {
	defer wg.Done()
	// read items from jobsCha
	for job := range jobsChan{
		processedItem := ProcessedItem{ID: job.ID, Result: fmt.Sprintf("Item %d Processed",job.ID)}
		
		// Simualte processing
		min := 100
		max := 900
		randomSleepTime := rand.Intn(max - min + 1) + min
		fmt.Printf("Routine %d will Process item %d\n", id, job.ID)
		fmt.Printf("I am %d. Sleeping for %v\n", job.ID, randomSleepTime)
		time.Sleep(time.Millisecond * time.Duration(randomSleepTime))


		//  send it to results
		resultsChan <- processedItem
	}
	
}


func main() {

	// Generate seed
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	numWorkers := 3
	numItems := 5
	jobsChan := make(chan Item, numItems)
	resultsChan := make(chan ProcessedItem)




	for i:=0; i< numWorkers; i++ {
		wg.Add(1)
		go worker(jobsChan, resultsChan, &wg, i)
	
	}

	
	go func(){
	// create 10 items, send each to jobsChan. close the channel
		for range numItems {
			id := 0
			data := ""
			fmt.Println("Enter item id and it's data: ")
			fmt.Scanln(&id)
			fmt.Scanln(&data)
			item := Item{ID: id, Data: data}
			jobsChan <- item
		}
		close(jobsChan)
	}()
	

	

	for range numItems{
		result := <- resultsChan
		fmt.Println("Results reached: ", result)
	}

	wg.Wait()
	close(resultsChan)



}
