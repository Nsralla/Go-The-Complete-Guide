package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/project/models"
	"example.com/project/utils"
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
    // User must be authenticated to create an event
    token := context.GetHeader("Authorization")
    if token == "" {
        context.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
        return
    }
    
    claims, err := utils.ValidateJWT(token)
    if err != nil {
        context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }
    
	// userEmail := claims.Email
	userID := claims.ID
    
    // Log authenticated user info
    // fmt.Printf("Authenticated user: ID=%d, Email=%s\n", userID, userEmail)

    // Extract data from request
    var new_event models.Event
    err = context.ShouldBindJSON(&new_event)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Use the actual user ID from JWT claims
    new_event.UserID = int(userID) // Convert uint to int if needed
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
	_,err = models.GetEvent(id_int)
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

	event.ID = id_int
	event.UserID = 1 // Assuming user ID is 1 for now
	err = event.Update()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error while updating event: %v", err)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": event})

}

func deleteEventById(context * gin.Context) {
	id := context.Param("id")
	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":fmt.Sprintf("Invalid format %v", err)} )
	}

	event, err := models.GetEvent(id_int)
	if err != nil{
		context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Event with ID %d not found", id_int)})
		return 
	}

	err = event.Delete()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error while deleting event: %v", err)})
		return
	}


}