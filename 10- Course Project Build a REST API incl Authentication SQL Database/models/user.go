package models

import (
	"errors"

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

func (u User) ValidateCredentials() (User, error) {
	storedUser, err := GetUser(u)
	if err != nil {
		return User{}, err
	}
	if !utils.CheckPasswordHash(u.Password, storedUser.Password) {
		return User{}, errors.New("invalid credentials")
	}
	return storedUser, nil
}

func GetUser(u User) (User, error) {
	query := `
		SELECT id, email, password
		FROM users
		WHERE email = ?
	`

	row := db.DB.QueryRow(query, u.Email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	// user.Password is the hashed password from DB
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func GetAllUsers() ([]User, error) {
	query := `
		SELECT id, email
		FROM users
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
