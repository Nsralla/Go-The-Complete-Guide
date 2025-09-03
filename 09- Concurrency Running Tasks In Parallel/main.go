package main
// Go routines doesn't return values, instead use channels to communicate
import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true // Signal that the task is done.
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // Signal that the task is done.
	close(doneChan) // Close the channel after all tasks are done. (Can be done only if we know that this is the longest task)
}

func main() {
	doneChan := make(chan bool) // Create a channel to signal task completion.
	go greet("Nice to meet you!", doneChan)
	go greet("How are you?", doneChan)
	go slowGreet("How ... are ... you ...?", doneChan)
	go greet("I hope you're liking the course!", doneChan)
	// Using just go will not wait for the goroutines to finish. so we need to wait for all goroutines to finish.
	// We can do this by waiting for the doneChan to receive a value for each goroutine we started.
	for range doneChan {
		<-doneChan
	}
}
