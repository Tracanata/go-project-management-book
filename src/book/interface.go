package book

type BookRepository interface {
	GetAllBook() ([]Book, error)
	GetBookById(id int) (*Book, error)
	SaveBook(book Book) error
}

type BookService interface {
	GetAllBooks() ([]Book, error)
	GetBookById(id int) (*Book, error)
	AddBook(book Book) error
}
