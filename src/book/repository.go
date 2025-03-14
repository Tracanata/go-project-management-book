package book

import (
	"errors"

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
	result := r.db.Create(&book)
	return result.Error
}

func (r *bookRepository) DeleteBook(id int) error {
	result := r.db.Where("id = ?", id).Delete(&Book{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Book Not Found")
	}
	return nil
}

func (r *bookRepository) UpdateBook(book Book) error {
	result := r.db.Where("codeBook = ?", book).Save(&book)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}
