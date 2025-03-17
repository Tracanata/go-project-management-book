package book

import (
	"fmt"
	"math/rand/v2"
)

type bookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]Book, error) {
	return s.repo.GetAllBook()
}

func (s *bookService) GetBookByCodeBook(code string) (*Book, error) {
	return s.repo.GetBookByCodeBook(code)
}

func (s *bookService) AddBook(book Book) error {
	book.Code_Book = fmt.Sprintf("B0%03dK", rand.IntN(999)+1)
	return s.repo.SaveBook(book)
}

func (s *bookService) DeleteBook(code string) error {
	return s.repo.DeleteBook(code)
}

func (s *bookService) UpdateBook(book Book) error {
	return s.repo.UpdateBook(book)
}
