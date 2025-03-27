package borrowed

import (
	"errors"
	"go-project-management-book/src/book"
	"time"
)

type BorrowedService struct {
	bookRepo   book.BookRepository
	borrowRepo BorrowedRepository
}

func NewBorrowedService(bookRepo book.BookRepository, borrowRepo BorrowedRepository) *BorrowedService {
	return &BorrowedService{
		bookRepo:   bookRepo,
		borrowRepo: borrowRepo,
	}
}

func (svc *BorrowedService) BorrowBook(userId, bookId int, days int, status string) error {
	stock, err := svc.bookRepo.GetBookStock(bookId)
	if err != nil {
		return errors.New("gagal mengambil data buku")
	}

	if stock <= 0 {
		return errors.New("buku tidak tersedia")
	}

	err = svc.bookRepo.DecreaseBookStock(bookId)
	if err != nil {
		return errors.New("gagal update stock buku")
	}

	now := time.Now()
	borrowed := Borrowed{
		UserId:     userId,
		BookId:     bookId,
		BorrowDate: now,
		Status:     "borrowed",
	}
	err = svc.borrowRepo.SetBorrowBook(&borrowed)
	if err != nil {
		return errors.New("gagal mencatat peminjaman")
	}

	return nil
}
