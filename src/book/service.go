package book

import (
	"errors"
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

func (s *bookService) GetBookByCodeBook(code string) (*RespDataBook, error) {
	book, err := s.repo.GetBookByCodeBook(code)

	if err != nil {
		return nil, errors.New("data buku tidak ada")
	}

	resp := &RespDataBook{
		Code_Book:  book.Code_Book,
		Title_Book: book.Title_Book,
		Author:     book.Author,
		Stock:      book.Stock,
	}
	return resp, nil
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
