package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)
type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error){
	if text == ""  {
		return Todo{}, errors.New("title and content cannot be empty")
	}
	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) ToString()  {
	fmt.Printf("Text: %s\n --------------------", todo.Text)
}

func (todo Todo) Save() error{
	// each todo has it's own file
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling todo to JSON:", err)
		return err
	}

	err = os.WriteFile(fileName, json, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return err
	}

	return nil
}