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

func (r *bookRepository) GetBookByCodeBook(code string) (*Book, error) {
	var book Book
	result := r.db.Where("code_book = ?", code).First(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r *bookRepository) SaveBook(book Book) error {
	result := r.db.Create(&book)
	return result.Error
}

func (r *bookRepository) DeleteBook(code string) error {
	result := r.db.Where("code_book = ?", code).Delete(&Book{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Book Not Found")
	}
	return nil
}

func (r *bookRepository) UpdateBook(book Book) error {
	var checkBook Book
	err := r.db.Where("code_book = ?", book.Code_Book).First(&checkBook).Error
	if err != nil {
		return err
	}
	return r.db.Model(&Book{}).Where("code_book = ?", book.Code_Book).Updates(book).Error
}
