package repository

import (
	"github.com/adliih/demo-bookstore-api/db"
	"github.com/adliih/demo-bookstore-api/models"
)

type BookRepository interface {
	FindAllBooks() ([]models.Book, error)
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) FindAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := db.GetDB().Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
