package service

import (
	"github.com/adliih/demo-bookstore-api/models"
	"github.com/adliih/demo-bookstore-api/repository"
)

type BookService interface {
	FindAllBooks() ([]models.Book, error)
	CreateBook(input models.CreateBookInput) error
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{bookRepo: repo}
}

func (s *bookService) FindAllBooks() ([]models.Book, error) {
	return s.bookRepo.FindAllBooks()
}

func (s *bookService) CreateBook(input models.CreateBookInput) error {
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
		Genre:  input.Genre,
	}
	return s.bookRepo.CreateBook(book)
}
