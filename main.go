package main

import (
	"github.com/adliih/demo-bookstore-api/controller"
	"github.com/adliih/demo-bookstore-api/db"
	"github.com/adliih/demo-bookstore-api/repository"
	"github.com/adliih/demo-bookstore-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db.InitDB(sqlite.Open("sqlite.db"), &gorm.Config{})

	router := gin.Default()

	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	v1 := router.Group("/api/v1")
	v1.GET("/books", bookController.GetAllBooks)

	router.Run(":8080")
}
