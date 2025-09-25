package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data string
}

type TaskResult struct {
	ID     int
	Result string
}

func(task *Task) processTask( resultChan chan TaskResult, wg * sync.WaitGroup) {

	defer wg.Done()

	min := 100
	max := 700
	randomDelay := rand.Intn(max - min + 1) + min 

	// Simulate task
	time.Sleep(time.Millisecond * time.Duration(randomDelay))

	//generate random string
	letters := []rune("abcdefghijklmnopqrstuvwx")
	randomString := make([]rune, 7)
	for i := range randomString{
		randomString[i] = letters[rand.Intn(len(letters))]
	}

	//Write to channel
	resultChan <- TaskResult{ID:task.ID, Result: string(randomString)}
}

func main() {
	// Seed generation
	rand.Seed(time.Now().UnixNano()) 

	var wg sync.WaitGroup

	// Create a slice of Tasks
	tasks := make([]Task, 5)
	tasks[0] = Task{ID: 1, Data:"Clita elitr no diam vero clita diam consetetur amet. Magna."}
	tasks[1] = Task{ID: 2, Data:"Clita elitr no diam vero clita diam consetetur amet. Magna."}
	tasks[2] = Task{ID: 3, Data:"Clita elitr no diam vero clita diam consetetur amet. Magna."}
	tasks[3] = Task{ID: 4, Data:"Clita elitr no diam vero clita diam consetetur amet. Magna."}
	tasks[4] = Task{ID: 5, Data:"Clita elitr no diam vero clita diam consetetur amet. Magna."}

	// Create a TaskResult channel
	taskResultChannel := make(chan TaskResult, 5)

	for i := range tasks{
		wg.Add(1)
		go tasks[i].processTask(taskResultChannel, &wg)
	}

	
	wg.Wait()
	close(taskResultChannel)
	
	for i:=0; i < len(tasks); i++{
		taskResult := <- taskResultChannel
		fmt.Printf("Task id: %d, Task result: %v \n",taskResult.ID, taskResult.Result)

	}
	
}