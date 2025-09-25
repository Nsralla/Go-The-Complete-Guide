package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// - Define Task and TaskResult structs: Task with ID (int) and Data (string), and TaskResult with ID (int) and Result (string).
// - Implement processTask(task Task, resultChan chan TaskResult): Simulate processing with a random delay, generate a result string, and send a TaskResult to the channel.
// - In main, create a slice of Tasks and a TaskResult channel.
// - Launch a goroutine for each Task: Each goroutine should call processTask, passing the task and the result channel.
// - In main, receive results from the channel: Print the ID and Result of each completed task. Ensure the main function waits for all tasks to finish.

type TaskResult struct{
	ID int 
	Result string
}

type Task struct{
	ID int 
	Data string
}

func New(id int, data string)Task{
	return Task{
		ID: id,
		Data: data,
	}
}


func main(){

	// Seed
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	task1 := New(1, "aaaaaaaaaaaa")
	task2 := New(2, "Dolores duo sed vero clita at, stet lorem vero dolor.")
	task3 := New(3, "Mild busen der lieb gut ein, noch guten mich.")
	task4 := New(4, "Would so his his congealed almost or sea,.")
	task5 := New(5, "Sit dolor dolore amet aliquyam diam dolor.")

	numTasks := 5

	tasks := make([]Task, 0,numTasks)
	tasks = append(tasks, task1)
	tasks = append(tasks, task2)
	tasks = append(tasks, task3)
	tasks = append(tasks, task4)
	tasks = append(tasks, task5)

	resultsChan := make(chan TaskResult)

	for _, task := range tasks{
		wg.Add(1)
		go processTask(task, resultsChan, &wg)
	}


	for i:=0; i< numTasks; i++ {
		result := <- resultsChan
		fmt.Println(result)
	}

	wg.Wait()
	close(resultsChan)
}


func processTask(task Task, resultChan chan <- TaskResult, wg * sync.WaitGroup){
	defer wg.Done()
	// Simulate Delay
	min := 100
	max := 500
	randomSleepTime := rand.Intn(max - min + 1) + min
	time.Sleep(time.Millisecond * time.Duration(randomSleepTime))

	// Write result to channel
	resultString := task.Data + fmt.Sprintf("%d", randomSleepTime)
	resultChan <- TaskResult{ID: task.ID, Result: resultString}
}