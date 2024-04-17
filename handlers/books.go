package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pteus/book_rest_api/models"
)

func GetBooks(c *gin.Context) {
	books, err := models.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	bookId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse book id"})
		return
	}

	book, err := models.GetBookById(bookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func AddBook(c *gin.Context) {
	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = book.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create book"})
		return
	}

	c.JSON(http.StatusCreated, book)

}
