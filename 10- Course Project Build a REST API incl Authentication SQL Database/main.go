package main

import (
	"example.com/project/db"
	"github.com/gin-gonic/gin"
	"example.com/project/routes"
)
func main(){

	// Init the database connection
	db.InitDB()
	defer db.DB.Close() // Close the database connection when the main function ends


	// Create a server
	server := gin.Default()
	routes.RegisterRoutes(server)


	// Start listening on the server for upcoming requests
	server.Run(":8080") // localhost:8080
}

