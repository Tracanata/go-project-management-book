package borrowed

import (
	"time"
)

type Borrowed struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	BookId     int       `json:"book_id"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date,omitempty"`
	Status     string    `json:"role" binding:"required,oneof=returned borrowed"`
}
