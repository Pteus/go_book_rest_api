package main

import (
	"github.com/gin-gonic/gin"

	"github.com/pteus/book_rest_api/db"
	"github.com/pteus/book_rest_api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
