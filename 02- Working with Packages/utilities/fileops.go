package utilities
import (
	"fmt"
	"os"
	"errors"
	"strconv"
)

func WriteToFile(value float64, fileName string) error {
	// convert balance to string
	stringValue := fmt.Sprint(value)
	// store balance in file, but at first convert string to list of bytes, 0644 is the permission
	err := os.WriteFile(fileName, []byte(stringValue), 0644)
	if err != nil {
		return errors.New("FAILED TO WRITE BALANCE TO FILE")
	}
	return nil
}

func ReadFromFile(fileName string) (float64, error){
	//read balance from file as bytes
	data, err :=os.ReadFile(fileName)
	// if an error occurred
	if err != nil {
		return 0, errors.New("FAILED TO READ BALANCE FILE")
	}
	// convert data to string
	stringValue := string(data)
	// convert string to float64
	value, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return 0, errors.New("FAILED TO PARSE BALANCE")
	}
	return value, nil

}