package controller

import (
	"strconv"

	"github.com/adliih/demo-bookstore-api/models"
	"github.com/adliih/demo-bookstore-api/service"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{bookService: bookService}
}

func (ctrl *BookController) GetAllBooks(c *gin.Context) {
	books, err := ctrl.bookService.FindAllBooks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, books)
}

func (ctrl *BookController) CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	if err := ctrl.bookService.CreateBook(input); err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(201, gin.H{"message": "Book created successfully"})
}

func (ctrl *BookController) UpdateBook(c *gin.Context) {
	bookIDStr := c.Param("bookID")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid book ID"})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	err = ctrl.bookService.UpdateBook(uint(bookID), input)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Book updated successfully"})
}

func (ctrl *BookController) DeleteBook(c *gin.Context) {
	bookIDStr := c.Param("bookID")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid book ID"})
		return
	}

	err = ctrl.bookService.DeleteBook(uint(bookID))
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Book deleted successfully"})
}
