package book

import (
	"errors"

	"gorm.io/gorm"
)

type IBookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &IBookRepository{db: db}
}

func (r *IBookRepository) GetAllBook() ([]Book, error) {
	var books []Book
	result := r.db.Find(&books)
	return books, result.Error
}

func (r *IBookRepository) GetBookByCodeBook(code string) (*Book, error) {
	var book Book
	result := r.db.Where("code_book = ?", code).First(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r *IBookRepository) SaveBook(book Book) error {
	result := r.db.Create(&book)
	return result.Error
}

func (r *IBookRepository) DeleteBook(code string) error {
	result := r.db.Where("code_book = ?", code).Delete(&Book{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Book Not Found")
	}
	return nil
}

func (r *IBookRepository) UpdateBook(book Book) error {
	var checkBook Book
	err := r.db.Where("code_book = ?", book.Code_Book).First(&checkBook).Error
	if err != nil {
		return err
	}
	return r.db.Model(&Book{}).Where("code_book = ?", book.Code_Book).Updates(book).Error
}

func (r *IBookRepository) GetBookStock(id int) (int, error) {
	var book Book
	err := r.db.Select("stock").Where("id = ?", id).First(&book).Error
	if err != nil {
		return 0, err
	}
	return book.Stock, nil
}

func (r *IBookRepository) DecreaseBookStock(id int) error {
	return r.db.Model(&Book{}).Where("id = ? AND stock > 0", id).UpdateColumn("stock", gorm.Expr("stock - 1")).Error
}
