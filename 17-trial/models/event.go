package models

import (
	"time"

	"example.com/app/db"
)

type Event struct {
	ID          int64
	Name        string    `bindin:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}



func (e Event) Create()(Event, error){

	query := `
		INSERT INTO events (name, description, location, dateTime, user_id) values(?,?,?,?,?)
	`
	newEvent, err := db.DB.Exec(query,e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil{
		return Event{}, err
	}
	id, err := newEvent.LastInsertId()
	if err != nil{
		return Event{}, nil
	}
	e.ID = id
	return e, nil

}
