package repository

import (
	"books/models"
	"database/sql"
	"errors"
	"fmt"
)

type BookRepository interface {
	CreateBook(book *models.Books) error
	GetBookByID(id int) (*models.Books, error)
	GetAllBooks() ([]models.Books, error)
	UpdateBook(book *models.Books) error
	DeleteBook(id int) error
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) CreateBook(book *models.Books) error {
	_, err := r.db.Exec("INSERT INTO books (name, author_name, price, available, issued, publisher, published_year, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		book.Name, book.AuthorName, book.Price, book.Available, book.Issued, book.Publisher, book.PublishedYear, book.Description)
	return err
}

func (r *bookRepository) GetBookByID(id int) (*models.Books, error) {
	row := r.db.QueryRow("SELECT id, name, author_name, price, available, issued, publisher, published_year, description FROM books WHERE id=?", id)
	var book models.Books
	err := row.Scan(&book.ID, &book.Name, &book.AuthorName, &book.Price, &book.Available, &book.Issued, &book.Publisher, &book.PublishedYear, &book.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) GetAllBooks() ([]models.Books, error) {
	var books []models.Books

	rows, err := r.db.Query("SELECT id, name, author_name, price, available, issued, publisher, published_year, description FROM books")
	if err != nil {
		return nil, fmt.Errorf("database query failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Books
		if err := rows.Scan(&book.ID, &book.Name, &book.AuthorName, &book.Price, &book.Available, &book.Issued, &book.Publisher, &book.PublishedYear, &book.Description); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return books, nil
}

func (r *bookRepository) UpdateBook(book *models.Books) error {
	result, err := r.db.Exec("UPDATE books SET name=?, author_name=?, price=?, available=?, issued=?, publisher=?, published_year=?, description=? WHERE id=?",
		book.Name, book.AuthorName, book.Price, book.Available, book.Issued, book.Publisher, book.PublishedYear, book.Description, book.ID)
	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error retrieving rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (r *bookRepository) DeleteBook(id int) error {
	result, err := r.db.Exec("DELETE FROM books WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error retrieving rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("book not found or no row deleted")
	}

	return nil
}
