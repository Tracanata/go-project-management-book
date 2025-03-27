package borrowed

type ReqBorrowedBook struct {
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
}
