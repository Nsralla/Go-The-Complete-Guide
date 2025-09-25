package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Implement a Go program that:
// 1- Defines a Worker struct with a Name and a method DoWork() int that simulates doing work by:
// 		.Sleeping for a random duration between 100ms and 700ms.
// 		.Returning a random result (int between 1–100).
// 2- Creates 5 workers: "Worker-1" to "Worker-5".
// 3- Launches each worker’s DoWork method concurrently using goroutines.
// 4- Uses channels to collect results.
// 5- Implements a timeout of 600ms for each worker. If a worker doesn't finish in time, print "Worker-X: timed out".
// 6- At the end, print: "Worker-X: timed out"
// 		.Each successful worker result.
// 		.The sum of all results that were returned on time.

type ResultsChan struct{
	WorkerName string 
	Result int
}
type Worker struct {
	Name string
}
func New( name string)Worker{
	return Worker{Name: name}
}
func (w * Worker) DoWork(wg *sync.WaitGroup, resultsChan chan <- ResultsChan){
	defer wg.Done()
	min := 100
	max := 700
	randomSleepTime := rand.Intn(max - min + 1) + min
	fmt.Printf("Sleep time for %v = %v\n",w.Name,randomSleepTime)


	min = 1
	max = 100
	randomResult := rand.Intn(max - min + 1) + min
	
	select {
		case <- time.After(time.Millisecond * 600):
			fmt.Printf("Worker %s didn't finish in time.\n",w.Name)
			resultsChan <- ResultsChan{WorkerName: w.Name, Result: -1}
			return
		case <- time.After(time.Millisecond * time.Duration(randomSleepTime)):
			resultsChan <- ResultsChan{WorkerName: w.Name, Result: randomResult}
	}
}


func main(){

	var wg sync.WaitGroup


	// Seed
	rand.Seed(time.Now().UnixNano())

	

	w1 := New("nsralla")
	w2 := New("ali")
	w3 := New("wessam")
	w4 := New("momen")
	w5 := New("wooooooooooo")

	workers := make([]Worker, 0,5)
	workers = append(workers, w1)
	workers = append(workers, w2)
	workers = append(workers, w3)
	workers = append(workers, w4)
	workers = append(workers, w5)

	resultsChan := make(chan ResultsChan,5)
	

	for _, worker := range workers{
		wg.Add(1)
		go func(){
			worker.DoWork(&wg, resultsChan)
		}()
	}

	wg.Wait()


	total :=0
	for i:=0; i<len(workers);i++{
		result := <- resultsChan
		if result.Result == -1{
			fmt.Printf("Worker %v: Timedout\n",result.WorkerName)
			continue
		}
		fmt.Printf("Worker %v Finished on time with res = %v \n",result.WorkerName, result.Result)
		total += result.Result
	}

	fmt.Printf("Final result: %v\n", total)
	

}