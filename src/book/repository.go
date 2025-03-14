package book

import (
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAllBook() ([]Book, error) {
	var books []Book
	result := r.db.Find(&books)
	return books, result.Error
}

func (r *bookRepository) GetBookById(id int) (*Book, error) {
	var book Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (r *bookRepository) SaveBook(book Book) error {
	return r.db.Create(&book).Error
}
