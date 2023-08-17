package controller

import (
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
