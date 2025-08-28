package main

import (
	"fmt"
	"example.com/structs/user"
	// "time"
)


func main(){
	firstName := getUserInput("Please enter your first name: ")
	lastName := getUserInput("Please enter your last name: ")
	birthDate := getUserInput("Please enter your birth date MM/DD/YYYY: ")
	
	// user := user.User{
	// 	firstName: firstName, // this will cause an error, because i made that this user parameters are not available outside their struct (private access modifier)
	// 	lastName:  lastName,
	// 	birthDate: birthDate,
	// 	createdAt: time.Now(),
	// }
	// USING User built in methods
	// user.PrintUserInfo()
	// // UPdating user info from built in method
	// user.UpdateLastName("Smith")
	// // USING User built in methods
	// user.PrintUserInfo()



	// or
	// CREATING A USER USING CONSTUCTOR METHOD
	newUser, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	newUser.PrintUserInfo()



	// create new admin
	admin, err := user.NewAdmin("admin@example.com", "password123", firstName, lastName, birthDate)
	if err != nil {
		fmt.Println("Error creating admin:", err)
		return
	}
	// admin.PrintUserInfo()
	admin.UpdateLastName("AdminLastName")
	admin.PrintUserInfo()
	// admin.User.PrintUserInfo()
}

func getUserInput(prompt string)(value string) {
	fmt.Print(prompt)
	// fmt.Scan(&value)
	fmt.Scanln(&value) // Entering a new line will end input
	return
}