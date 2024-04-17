package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all books"})
}
