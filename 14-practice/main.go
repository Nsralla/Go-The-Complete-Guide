package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RawData struct{
	ID int
	Value int
}

type ProcessedData struct{
	ID int 
	TransformedValue int
}

func (rawData * RawData) transformData(outputChan chan <- ProcessedData, wg * sync.WaitGroup){

	defer wg.Done()

	// Random delay
	min := 100
	max := 900
	randomNumber := rand.Intn(max - min + 1) + min
	fmt.Printf("I am %d, sleeping for %d\n", rawData.ID, randomNumber)
	time.Sleep(time.Millisecond * time.Duration(randomNumber))

	// Simulate transformation
	tv := rawData.Value * 2

	outputChan <- ProcessedData{ID: rawData.ID, TransformedValue:tv}
}

func main(){
	// Create Seed
	rand.Seed(time.Now().UnixNano())

	// Create wait group
	var wg sync.WaitGroup

	//define length of raw data
	rawDataLen := 5

	// Create slice
	rawData := make([]RawData, rawDataLen)

	
	// read raw data from terminal
	id := 0
	value := 0
	for i := range rawData{
		fmt.Println("Enter raw data id and it's value: ")
		fmt.Scanln(&id)
		fmt.Scanln(&value)
		rawData[i] = RawData{ID: id, Value: value}
	}

	// Create channel for Transformed data
	pdChannel := make(chan ProcessedData, rawDataLen)

	// Call go routines
	for i:=0; i< rawDataLen; i++{
		wg.Add(1)
		go rawData[i].transformData(pdChannel, &wg) 
	}

	wg.Wait()
	close(pdChannel)

	sum := 0
	for range rawData{
		processdDate := <- pdChannel
		fmt.Printf("I am %d , tv =  %d\n", processdDate.ID, processdDate.TransformedValue)
		sum += processdDate.TransformedValue
	}

	fmt.Println("Sum = ",sum)

}