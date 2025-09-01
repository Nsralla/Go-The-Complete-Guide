package filemanager

import (
	"os"
	"bufio"
	"fmt"
)
func ReadLinesFromFile(filePath string)([]string, error){
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		file.Close()
		return nil, err
	}

	// Read the lines from the file
	lines := []string{}
	reader := bufio.NewScanner(file)
	for reader.Scan() { // this will read line by line
		line := reader.Text()
		lines = append(lines, line)
	}
	// Check for errors
	err = reader.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	// Close file
	file.Close()
	return lines, nil

}