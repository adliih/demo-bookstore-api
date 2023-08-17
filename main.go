package main

import (
	"github.com/adliih/demo-bookstore-api/controller"
	"github.com/adliih/demo-bookstore-api/db"
	_ "github.com/adliih/demo-bookstore-api/docs"
	"github.com/adliih/demo-bookstore-api/repository"
	"github.com/adliih/demo-bookstore-api/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	v1.POST("/books", bookController.CreateBook)
	v1.PUT("/books/:bookID", bookController.UpdateBook)
	v1.DELETE("/books/:bookID", bookController.DeleteBook)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
