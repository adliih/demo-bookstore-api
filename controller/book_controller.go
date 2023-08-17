package controller

import (
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