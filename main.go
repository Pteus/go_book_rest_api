package main

import (
	"github.com/gin-gonic/gin"

	"github.com/pteus/book_rest_api/routes"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
