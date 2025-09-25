package main

import (
	"fmt"
	"sync"
)

// Question: Build a Concurrent User Logger

// Write a Go program that:
//     1.  Defines a User struct with fields: ID (string), Name (string).
//     2.  Defines a Logger interface with the following methods
// 					. LogUser(user User).
// 					. GetUser(id string) User
// 					. GetAllUsers()[]User
//     3.  Implements a MemoryLogger struct that:
//         •   Stores logged users in a map[string]User, keyed by ID.
//         •   Is safe for concurrent access.
//     4.  Spawns 5 goroutines, each logging a different User via MemoryLogger.
//     5.  After all goroutines finish, prints all logged users from the map.
// Note: Map datatype is not thread safe.
type Logger interface{
	LogUser(user User)
	GetUser(id string) User
	GetAllUsers()[]User
}
type MemoryLogger struct{
	LoggedUsers map[string]User 
	Mu sync.RWMutex
}
func(ml *MemoryLogger) LogUser(user User, wg * sync.WaitGroup){
	defer wg.Done()
	ml.Mu.Lock()
	defer ml.Mu.Unlock()
	ml.LoggedUsers[user.ID] = user
}
func (ml *MemoryLogger) GetUser(id string)User{
	ml.Mu.RLock()
	defer ml.Mu.RUnlock()
	return ml.LoggedUsers[id]
}
func (ml *MemoryLogger) GetAllUsers()[]User{
	ml.Mu.RLock()
	defer ml.Mu.RUnlock()
	var users []User
	for key,_ := range ml.LoggedUsers{
		users = append(users, ml.LoggedUsers[key])
	}
	return users
}

func New(id, name string)User{
	return User{
		ID: id,
		Name: name,
	}
}



type User struct{
	ID string 
	Name string
}
func main(){
	var wg sync.WaitGroup
	user1 := New("1", "nsralla")
	user2 := New("2", "ali")
	user3 := New("3", "loo")
	user4 := New("4", "zidan")
	user5 := New("5", "ibrahim")
	users := make([]User, 5)
	users[0] = user1
	users[1] = user2
	users[2] = user3
	users[3] = user4
	users[4] = user5

	memoryLogger := MemoryLogger{
		LoggedUsers: map[string]User{},
		Mu: sync.RWMutex{},
	}
	

	numRoutines := 5
	for i:=0; i< numRoutines; i++{
		wg.Add(1)
		go memoryLogger.LogUser(users[i], &wg)
	}

	wg.Wait()

	loggedUsers := memoryLogger.GetAllUsers()
	for _,user := range loggedUsers{
		fmt.Println(user)
	}


}