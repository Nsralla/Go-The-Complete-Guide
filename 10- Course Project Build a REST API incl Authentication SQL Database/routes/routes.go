package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Register a handler for "/events"
	server.GET("/events/:id", getEventByID)

	server.PUT("/events/:id", updateEventByID) // User must be authenticated to update an event
	server.DELETE("/events/:id", deleteEventById) // User must be authenticated to delete an event
	server.POST("/events", createEvent) // User must be authenticated to create an event
	
	server.GET("/events", getEvents) 

	server.POST("/signup", createUser)
	server.POST("/login", login)

	server.GET("/users", getAllUsers) // For testing purposes only
	
}
