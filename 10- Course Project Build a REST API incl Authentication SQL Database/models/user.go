package models

import (
	"example.com/project/db"
	"example.com/project/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

func (u User) Save() (User, error) {
	query := `
		INSERT INTO USERS (email, password) VALUES (?, ?)
	`

	// Hash the password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return User{}, err
	}

	result, err := db.DB.Exec(query, u.Email, hashedPassword)
	if err != nil {
		return User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}
	u.ID = id
	return u, nil
}
