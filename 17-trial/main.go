package main

import (
	"example.com/app/routes"
	"example.com/app/db"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
