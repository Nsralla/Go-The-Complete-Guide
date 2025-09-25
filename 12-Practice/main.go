package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Worker struct {
	Name string
}

type ResultChan struct{
	Idx int
	Result int
}

func New(name string)Worker{
	return Worker{
		Name: name,
	}
}

func (w *Worker) DoWork(wg * sync.WaitGroup, ch chan int, done chan struct{}) {
	defer wg.Done()
	// Generate a random number to sleep
	min := 100
	max := 700
	randomSleepTime := rand.Intn(max - min + 1) + min
	fmt.Printf("I am worker %s, I will sleep for %d\n", w.Name, randomSleepTime)

	select{
		case <- time.After(time.Millisecond * time.Duration(randomSleepTime)):
				// Generate random result
				min = 1
				max = 100
				result := rand.Intn(max - min + 1) + min 

				select {
					case ch <- result:
						close(ch)

					case <-done:
						return
				}
		case <-done:
			return 
	}	
}

func main() {

	rand.Seed(time.Now().UnixNano())

	// Generate workers
	numWorkers := 5
    workers := make([]Worker, numWorkers)
    worker1 := New("nsr")
    worker2 := New("mo")
    worker3 := New("loo")
    worker4 := New("noo")
    worker5 := New("roo")

    workers[0] = worker1
    workers[1] = worker2
    workers[2] = worker3
    workers[3] = worker4
    workers[4] = worker5

	// Create wg
	var wg sync.WaitGroup

	resultsChan := make(chan ResultChan)

	for i, worker := range workers{
		wg.Add(1)
		i := i
		worker := worker
		go func(){
				ch:= make(chan int, 1)
				done := make(chan struct{})
				go worker.DoWork(&wg,ch, done)
				select{
					case result := <- ch: // In case all went well
						resultsChan <-  ResultChan{
							Idx: i,
							Result: result,
						}
					
					case <- time.After(time.Millisecond * 600): // In case time finishes before
						fmt.Printf("I am %s worker. Timed out.", worker.Name)
						close(done)
						resultsChan <- ResultChan{
							Idx: i,
							Result: -1,
						}	
				}
		}()
		
	}

	
	
	total := 0
	for i:=0; i < len(workers); i++{
		result := <- resultsChan
		if result.Result !=-1 {
			fmt.Printf("Woker %s result: %d\n", workers[result.Idx], result.Result)
			total += result.Result
		}
	}
	
	fmt.Println("Total sum : ",total)


	wg.Wait()

}