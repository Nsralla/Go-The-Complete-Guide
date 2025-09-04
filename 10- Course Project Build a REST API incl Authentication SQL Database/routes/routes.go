package routes

import (
	"example.com/project/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Register a handler for "/events"
	server.GET("/events/:id", getEventByID)


	server.PUT("/events/:id", middleware.Authenticate,updateEventByID) // User must be authenticated to update an event
	// so now any request to add a new event will first go through the Authenticate middleware, validate if the use is logged in
	// and then only allow the request to proceed to createEvent handler if the user is authenticated. also if the user is authinticated, the user id 
	// will be attached to the context which we can use in createEvent to associate the event with the user creating it.
	// and to reach the id you have to use context.GetInt64("userID") which will return the user id stored in the context by the middleware
	server.DELETE("/events/:id",middleware.Authenticate ,deleteEventById) // User must be authenticated to delete an event
	server.POST("/events", middleware.Authenticate ,createEvent) // User must be authenticated to create an event
	
	
	server.GET("/events", getEvents) 

	server.POST("/signup", createUser)
	server.POST("/login", login)

	server.GET("/users", getAllUsers) // For testing purposes only
	
}
