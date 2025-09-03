package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"encoding/json"
	"time"
)
type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath string, outputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fm *FileManager) ReadLinesFromFile()(lines []string, err error){
	// Open the file
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		file.Close()
		return nil, err
	}

	// Read the lines from the file
	lines = []string{}
	reader := bufio.NewScanner(file)
	for reader.Scan() { // this will read line by line
		line := reader.Text()
		lines = append(lines, line)
	}
	// Check for errors
	err = reader.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, errors.New("error reading file: " + err.Error())
	}
	// Close file
	file.Close()
	return lines, nil

}

func (fm *FileManager) WriteJson(data any) error {
	// Create or truncate the file
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return errors.New("error creating file: " + err.Error())
	}

	time.Sleep(3 * time.Second) // Simulate a long write operation

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return errors.New("error encoding JSON: " + err.Error())
	}
	file.Close()
	return nil
}