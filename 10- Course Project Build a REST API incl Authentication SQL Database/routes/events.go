package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/project/models"
	"github.com/gin-gonic/gin"
)

func getEventByID(context *gin.Context) {
	fmt.Printf("Type of context.Param(\"id\"): %T\n", context.Param("id"))
	// Extract the ID from the URL
	id := context.Param("id")
	// Convert the ID to an integer 64
	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a number"})
		return
	}
	// Fetch the event from the database
	event, err := models.GetEvent(id_int)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Event with ID %d not found", id_int)})
		return
	}
	// Return the event as JSON
	context.JSON(http.StatusOK, event)
}

// By default it takes `context`
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println("Error while fetching events:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while fetching events: %v", err)})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// Extract data from request.
	var new_event models.Event
	err := context.ShouldBindJSON(&new_event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	new_event.UserID = 1
	new_event, err = new_event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while saving event: %v", err)})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created successfully", "event": new_event})
}
