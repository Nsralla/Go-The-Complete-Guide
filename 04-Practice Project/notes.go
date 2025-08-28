package main

import (
	"bufio"
	"fmt"
	"notes-app/note"
	"os"
	"strings"
)


func main(){
	title, content := getUserInput()
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	newNote.ToString()
	err = newNote.SaveToFile()
	if err != nil {
		fmt.Println("Error saving note to file:", err)
		return
	}
}

func getUserInput()(string, string){
	var title, content string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter note title: ")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Print("Enter note content: ")
	content, _ = reader.ReadString('\n') // till what to keep reading? until \n
	content = strings.TrimSpace(content) // remove leading and trailing whitespace

	return title, content
}