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
	// Get the user ID from the context (set by the authentication middleware)
	userID := context.GetInt64("userID")
	fmt.Println("FROM evenets, Authenticated user ID from context:", userID)

    // Extract data from request
    var new_event models.Event
    err := context.ShouldBindJSON(&new_event)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Use the actual user ID from JWT claims
    new_event.UserID = int64(userID) // Convert uint to int if needed
    new_event, err = new_event.Save()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while saving event: %v", err)})
        return
    }
    
    context.JSON(http.StatusCreated, gin.H{"message": "Event Created successfully", "event": new_event})
}


func updateEventByID(context * gin.Context){
	// Extract the ID from the URL
	id := context.Param("id")
	// Convert the ID to an integer 64
	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid Id %s", id)})
		return
	}

	// Verify if the event with the given ID exists
	old_event,err := models.GetEvent(id_int)
	if err != nil{
		context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Event with ID %d not found", id_int)})
		return
	}

	// Store the updated event at variable event
	var event models.Event
	err = context.ShouldBind(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid event data: %v", err)})
		return
	}


	// User id who created the event
	user_who_created_event := old_event.UserID
	fmt.Println("User who created the event ID:", user_who_created_event)
	// user who made the request
	user_who_made_request := context.GetInt64("userID")
	fmt.Println("User who made the request ID:", user_who_made_request)
	if user_who_created_event != user_who_made_request {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to update this event"})
		return
	}

	event.ID = id_int
	event.UserID = user_who_made_request 
	err = event.Update()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error while updating event: %v", err)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": event})

}

func deleteEventById(context * gin.Context) {
	eventID := context.Param("id") // Extract the ID from the URL
	event_id_int, err := strconv.ParseInt(eventID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid format %v", err)})
		return
	}
	

	event, err := models.GetEvent(event_id_int)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Event with ID %d not found", event_id_int)})
		return
	}

	// Extract user id who made this request
	user_sent_request := context.GetInt64("userID")
	// Extract user id who created this event
	user_created_event_id := event.UserID
	if user_sent_request != user_created_event_id {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error while deleting event: %v", err)})
		return
	}


}


