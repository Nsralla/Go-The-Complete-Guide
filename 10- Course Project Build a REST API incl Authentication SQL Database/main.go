package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example.com/project/models"
)
func main(){
	// Create a server
	server := gin.Default()
	// Register a handler for "/events"
	server.GET("/events", eventsGetHandler) 
	server.POST("/events", createEvent)

	// Start listening on the server for upcoming requests
	server.Run(":8080") // localhost:8080
}

// By defualt it takes `context`
func eventsGetHandler(context *gin.Context){
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}


func createEvent(context *gin.Context){
	// Extract data from request.
	var new_event models.Event
	err := context.ShouldBindJSON(&new_event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	new_event.ID = 1
	new_event.UserID = 1
	new_event.Save()
	context.JSON(http.StatusCreated, gin.H{"message":"Event Created successfully", "event":new_event})
}

