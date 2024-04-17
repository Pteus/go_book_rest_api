package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/pteus/book_rest_api/handlers"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/books", handlers.GetBooks)
	server.GET("/books/:id", handlers.GetBook)
	server.POST("/books", handlers.AddBook)
}
