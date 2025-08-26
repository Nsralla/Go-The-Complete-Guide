package main

import (
	"fmt"
	"time"
)

//  Struct: Groups related data into one location
type User struct {
	FirstName string
	LastName string
	BirthDate string 
	createdAt time.Time
}
func main(){
	firstName := getUserInput("Please enter your first name: ")
	lastName := getUserInput("Please enter your last name: ")
	birthDate := getUserInput("Please enter your birth date MM/DD/YYYY: ")
	
	user := User{
		FirstName: firstName,
		LastName: lastName,
		BirthDate: birthDate,
		createdAt: time.Now(),
	}

	printUserInfo(&user)
}

func printUserInfo(user *User) {
	fmt.Println("User Information:")
	fmt.Println("First Name:", user.FirstName)
	// or
	fmt.Println("First Name:", (*user).FirstName) // same as above
	fmt.Println("Last Name:", user.LastName)
	fmt.Println("Birth Date:", user.BirthDate)
	fmt.Println("Created At:", user.createdAt)
}

func getUserInput(prompt string)(value string) {
	fmt.Print(prompt)
	fmt.Scan(&value)
	return
}