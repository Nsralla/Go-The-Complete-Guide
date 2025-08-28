package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)
type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error){
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}
	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) ToString()  {
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n --------------------", note.Title, note.Content, note.CreatedAt.Format(time.RFC1123))
}

func (note Note) SaveToFile() error{
	// each note has it's own file
	fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		fmt.Println("Error marshalling note to JSON:", err)
		return err
	}

	err = os.WriteFile(fileName, json, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return err
	}

	return nil
}