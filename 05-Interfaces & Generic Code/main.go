package main

import (
	"bufio"
	"fmt"
	"notes-app/note"
	"notes-app/todo"
	"os"
	"strings"
)

type Saver interface {
	Save() error
	// ToString() 
}
// this tells: any struct of type interface must have a Save() method that returns an error

// or you can embed the interface
// type Display interface {
// 	ToString()
// }
type Outputter interface {
	Saver
	ToString()
}

func main(){

	

	// register new note
	title, content := getUserInput()
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	// register new todo
	newTodo, err := todo.New(content)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	err = outputData(newNote)
	if err != nil {
		fmt.Println("Error outputting note:", err)
		return
	}
	err = outputData(newTodo)
	if err != nil {
		fmt.Println("Error outputting todo:", err)
		return
	}

	// save user todo
	// newTodo.ToString()
	// err = saveUserData(newTodo)
	// if err != nil {
	// 	fmt.Println("Error saving todo:", err)
	// 	return
	// }

	// save user note
	// newNote.ToString()
	// err = saveUserData(newNote)
	// if err != nil {
	// 	fmt.Println("Error saving note:", err)
	// 	return
	// }

	printAny("Hello, World!")
	printAny(42)
	printAny(newNote)
	printAny(newTodo)
}

func printAny(value any) {
	fmt.Println(value)
}
func outputData(data Outputter) error{
	data.ToString() // this struct must implement the ToString method
	err := saveUserData(data) 
	if err != nil {
		fmt.Println("Error saving data:", err)
		return err
	}
	return nil
}



func saveUserData(userData Saver) error{
	err := userData.Save() // this struct must implement the Save method
	if err != nil {
		fmt.Println("Error saving note to file:", err)
		return err
	}
	return nil
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