package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events",handleGetEvents)
	server.POST("/events",handleAddEvent )
	server.GET("/events/:id", handleGetEvent)
}