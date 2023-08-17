package db

import (
	"github.com/adliih/demo-bookstore-api/models"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// InitDB initializes the database connection.
func InitDB(dialector gorm.Dialector, opts ...gorm.Option) {
	var err error
	DB, err = gorm.Open(dialector, opts...)
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate your models to the database
	DB.AutoMigrate(&models.Book{})
}

// GetDB returns the database instance.
func GetDB() *gorm.DB {
	return DB
}
