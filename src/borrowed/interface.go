package borrowed

type BorrowedRepository interface {
	SetBorrowBook(borrowed *Borrowed) error
}

type IService interface {
	BorrowBook(userId, bookId int, days int, status string) error
}
