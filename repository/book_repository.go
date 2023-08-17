package repository

import (
	"github.com/adliih/demo-bookstore-api/db"
	"github.com/adliih/demo-bookstore-api/models"
)

type BookRepository interface {
	FindAllBooks() ([]models.Book, error)
	CreateBook(book models.Book) error
	UpdateBook(bookID uint, updatedBook models.Book) error
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

func (r *bookRepository) CreateBook(book models.Book) error {
	return db.GetDB().Create(&book).Error
}

func (r *bookRepository) UpdateBook(bookID uint, updatedBook models.Book) error {
	return db.GetDB().Model(&models.Book{}).
		Where(&models.Book{ID: bookID}).
		Updates(updatedBook).Error
}
