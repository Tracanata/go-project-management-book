package book

type bookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]Book, error) {
	return s.repo.GetAllBook()
}

func (s *bookService) GetBookById(id int) (*Book, error) {
	return s.repo.GetBookById(id)
}

func (s *bookService) AddBook(book Book) error {
	return s.repo.SaveBook(book)
}
