package book

type BookRepository interface {
	GetAllBook() ([]Book, error)
	GetBookByCodeBook(code string) (*Book, error)
	SaveBook(book Book) error
	DeleteBook(code string) error
	UpdateBook(book Book) error
	GetBookStock(id int) (int, error)
	DecreaseBookStock(id int) error
}

type BookService interface {
	GetAllBooks() ([]Book, error)
	GetBookByCodeBook(code string) (*RespDataBook, error)
	AddBook(book Book) error
	DeleteBook(code string) error
	UpdateBook(book Book) error
}
