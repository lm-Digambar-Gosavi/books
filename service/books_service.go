package service

import (
	"books/models"
	"books/repository"
	"errors"
	"log"
)

type BookService interface {
	CreateBook(book *models.Books) error
	GetBookByID(id int) (*models.Books, error)
	GetAllBooks() ([]models.Books, error)
	UpdateBook(book *models.Books) error
	DeleteBook(id int) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) CreateBook(book *models.Books) error {
	if book.Name == "" || book.AuthorName == "" || book.Price < 0 || book.Available < 0 || book.Issued < 0 {
		return errors.New("invalid book details")
	}
	err := s.repo.CreateBook(book)
	if err != nil {
		log.Println("Error Occurred:", err)
		return err
	}
	return nil
}

func (s *bookService) GetBookByID(id int) (*models.Books, error) {
	book, err := s.repo.GetBookByID(id)
	if err != nil {
		log.Println("Error Occurred:", err)
		return nil, err
	}
	return book, nil
}

func (s *bookService) GetAllBooks() ([]models.Books, error) {
	books, err := s.repo.GetAllBooks()
	if err != nil {
		log.Println("Error Occurred:", err)
		return nil, err
	}
	return books, nil
}


func (s *bookService) UpdateBook(book *models.Books) error {
	if book.Name == "" || book.AuthorName == "" || book.Price < 0 || book.Available < 0 || book.Issued < 0 {
		return errors.New("invalid book details")
	}
	err := s.repo.UpdateBook(book)
	if err != nil {
		log.Println("Error Occurred:", err)
		return err
	}
	return nil
}

func (s *bookService) DeleteBook(id int) error {
	err := s.repo.DeleteBook(id)
	if err != nil {
		log.Println("Error Occurred:", err)
		return err
	}
	return nil
}
