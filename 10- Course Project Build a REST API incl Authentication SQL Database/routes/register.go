package routes

import (
	"strconv"
"fmt"
	"example.com/project/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	// Get the user ID from the context (set by the authentication middleware)
	userID_who_made_request := context.GetInt64("userID")
	if userID_who_made_request == 0 {
		context.JSON(401, gin.H{"message": "You must be logged in to register for an event"})
		return
	}
	// Extract the event ID from the URL
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Invalid event ID"})
		return
	}
	// Check if the event exists
	event, err := models.GetEvent(eventID)
	if err != nil {
		context.JSON(404, gin.H{"message": "Event not found"})
		return
	}

	// Register the user for the event
	err = event.RegisterUserForEvent(userID_who_made_request)
	if err != nil {
		context.JSON(500, gin.H{"message": "Failed to register for event"})
		return
	}
	context.JSON(200, gin.H{"message": "Successfully registered for the event"})
	
}

func cancelRegistrationFromEvent(context *gin.Context) {
	// Find the user ID from the URL
	eventID := context.Param("id")
	eventID_int, err := strconv.ParseInt(eventID, 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEvent(eventID_int)
	if err != nil {
		context.JSON(404, gin.H{"message": "Event not found"})
		return
	}

	// FIND the user ID from the context (set by the authentication middleware)
	userID := context.GetInt64("userID")
	if userID == 0 {
		context.JSON(401, gin.H{"message": "You must be logged in to cancel registration"})
		return
	}

	// cancel the registration
	err = event.CancelUserRegistration(userID)
	if err != nil {
		context.JSON(500, gin.H{"message": "Failed to cancel registration"})
		return
	}
	context.JSON(200, gin.H{"message": "Registration cancelled successfully"})
}

func getAllRegistrations(context *gin.Context) {
	// Get the user ID from the context (set by the authentication middleware)
	userID := context.GetInt64("userID")
	if userID == 0 {
		context.JSON(401, gin.H{"message": "You must be logged in to view your registrations"})
		return
	}
	fmt.Println("User ID who is looking for registrations:", userID)

	// Fetch all registrations for the user
	registrations, err := models.GetEventRegistrationsByUserID(userID)
	if err != nil {
		context.JSON(500, gin.H{"message": "Failed to retrieve registrations", "error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"registrations": registrations})
}

