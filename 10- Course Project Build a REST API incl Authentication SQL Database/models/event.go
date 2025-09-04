package models

import (
	"time"
)

type Event struct {
	ID int
	Name string `binding:"required"` // must be attached with the post request coming from the client side
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int 

}

func New(id int, name string, description string, location string, dateTime time.Time, userID int) Event {
	return Event{
		ID:          id,
		Name:       name,
		Description: description,
		Location:   location,
		DateTime:  dateTime,
		UserID:    userID,
	}
}

var events = []Event{}

func (e Event) Save(){
	events = append(events, e)
}

func  GetAllEvents() []Event{
	return events
}