package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Register a handler for "/events"
	server.GET("/events/:id", getEventByID)
	server.PUT("/events/:id", updateEventByID)
	server.DELETE("/events/:id", deleteEventById)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.POST("/signup", createUser)
	
}
