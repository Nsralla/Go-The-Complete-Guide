package user
import (
	"fmt"
	"errors"
	"time"
)
//  Struct: Groups related data into one location
type User struct {
	firstName string //// if it named firtName it wouldn't be accessible outside this package, FirstName will do the opposite
	lastName  string
	birthDate string
	createdAt time.Time
}
//attach a method to the User struct
// HERE YOU ARE PASSING A COPY AND NOT THE ORIGINAL STRUCT.
func (user User) PrintUserInfo() {
	fmt.Println("User Information:")
	fmt.Println("First Name:", user.firstName)
	// or
	// fmt.Println("First Name:", (*user).FirstName) // same as above
	fmt.Println("Last Name:", user.lastName)
	fmt.Println("Birth Date:", user.birthDate)
	fmt.Println("Created At:", user.createdAt)
}
// IF YOU WANT TO CHANGE THE ORIGINAL STRUCT, YOU NEED TO USE A POINTER RECEIVER.
func (user *User) UpdateLastName(newLastName string){
	user.lastName = newLastName
	//or
	// (*user).LastName = newLastName // same as above
}

// CREATING CONSTRUCTOR FUNCTION, returning a pointer not a copy (you can use any you like)
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}