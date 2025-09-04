package models

import (
	"time"
	"example.com/project/db"
)

type Event struct {
	ID int64
	Name string `binding:"required"` // This means: Name must be attached with the post request coming from the client side
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int 

}

func New(id int64, name string, description string, location string, dateTime time.Time, userID int) Event {
	return Event{
		ID:          id,
		Name:       name,
		Description: description,
		Location:   location,
		DateTime:  dateTime,
		UserID:    userID,
	}
}


func (e Event) Save() (Event, error){
	insertEventQuery := `
		INSERT INTO events (name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)
	`
	result, err := db.DB.Exec(insertEventQuery, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	// Exec is used when you have a query that changes the database statement
	// Query is used when you want to fetch data from the database
	// Exec returns a Result type and an error
	if err != nil {
		return Event{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Event{}, err
	}
	e.ID = id
	return e, nil

}

func  GetAllEvents() ([]Event, error){
	getAllEventsQuery := `
		SELECT * 
		FROM events
	`
	rows, err := db.DB.Query(getAllEventsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next(){
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEvent(id int64)(Event, error){
	query := `
		SELECT *
		FROM events
		WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}