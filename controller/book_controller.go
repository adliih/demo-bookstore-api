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

// GetAllBooks returns a list of all books.
// @Summary Get all books
// @Description Get a list of all books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/books [get]
func (ctrl *BookController) GetAllBooks(c *gin.Context) {
	books, err := ctrl.bookService.FindAllBooks()
	if err != nil {
		c.JSON(500, &models.ErrorResponse{Error: "Internal server error"})
		return
	}
	c.JSON(200, books)
}

// CreateBook creates a new book.
// @Summary Create a new book
// @Description Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Param input body models.CreateBookInput true "Book data"
// @Success 201 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/books [post]
func (ctrl *BookController) CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, &models.ErrorResponse{Error: "Bad request"})
		return
	}

	if err := ctrl.bookService.CreateBook(input); err != nil {
		c.JSON(500, &models.ErrorResponse{Error: "Internal server error"})
		return
	}

	c.JSON(201, &models.SuccessResponse{Message: "Book created successfully"})
}

// UpdateBook updates an existing book.
// @Summary Update a book
// @Description Update an existing book
// @Tags books
// @Accept json
// @Produce json
// @Param bookID path int true "Book ID"
// @Param input body models.UpdateBookInput true "Updated book data"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/books/{bookID} [put]
func (ctrl *BookController) UpdateBook(c *gin.Context) {
	bookIDStr := c.Param("bookID")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(400, &models.ErrorResponse{Error: "Invalid book ID"})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, &models.ErrorResponse{Error: "Bad request"})
		return
	}

	err = ctrl.bookService.UpdateBook(uint(bookID), input)
	if err != nil {
		c.JSON(500, &models.ErrorResponse{Error: "Internal server error"})
		return
	}

	c.JSON(200, &models.SuccessResponse{Message: "Book updated successfully"})
}

// DeleteBook deletes a book by ID.
// @Summary Delete a book
// @Description Delete a book by ID
// @Tags books
// @Param bookID path int true "Book ID"
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/books/{bookID} [delete]
func (ctrl *BookController) DeleteBook(c *gin.Context) {
	bookIDStr := c.Param("bookID")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(400, &models.ErrorResponse{Error: "Invalid book ID"})
		return
	}

	err = ctrl.bookService.DeleteBook(uint(bookID))
	if err != nil {
		c.JSON(500, &models.ErrorResponse{Error: "Internal server error"})
		return
	}

	c.JSON(200, &models.SuccessResponse{Message: "Book deleted successfully"})
}
