package borrowed

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewBorrowedRepository(db *gorm.DB) BorrowedRepository {
	return &Repository{db: db}
}

func (repo *Repository) SetBorrowBook(borrowed *Borrowed) error {
	return repo.db.Create(&borrowed).Error
}
