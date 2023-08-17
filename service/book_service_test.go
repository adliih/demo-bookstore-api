package service

import (
	"testing"

	"github.com/adliih/demo-bookstore-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockedBookRepository struct {
	mock.Mock
}

func (m *mockedBookRepository) FindAllBooks() ([]models.Book, error) {
	args := m.Called()
	return args.Get(0).([]models.Book), args.Error(1)
}

func TestFindAllBooks(t *testing.T) {
	mockRepo := new(mockedBookRepository)
	mockBooks := []models.Book{
		{Title: "Book 1", Author: "Author 1"},
		{Title: "Book 2", Author: "Author 2"},
	}
	mockRepo.On("FindAllBooks").Return(mockBooks, nil)

	service := NewBookService(mockRepo)
	books, err := service.FindAllBooks()

	assert.NoError(t, err)
	assert.NotNil(t, books)
	assert.Len(t, books, len(mockBooks))

	mockRepo.AssertExpectations(t)
}
