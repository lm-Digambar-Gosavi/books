package handlers

import (
	"books/models"
	"books/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Books
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, `{"status_code": 400, "message": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if err := h.bookService.CreateBook(&book); err != nil {
		http.Error(w, `{"status_code": 500, "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book created successfully"})
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, `{"status_code": 400, "message": "Invalid ID"}`, http.StatusBadRequest)
		return
	}

	book, err := h.bookService.GetBookByID(id)
	if err != nil {
		http.Error(w, `{"status_code": 404, "message": "Book not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		http.Error(w, `{"status_code": 500, "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, `{"status_code": 400, "message": "Invalid book ID"}`, http.StatusBadRequest)
		return
	}

	var book models.Books
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, `{"status_code": 400, "message": "Invalid request body"}`, http.StatusBadRequest)
		return
	}
	book.ID = id
	// Call service layer to update book
	if err := h.bookService.UpdateBook(&book); err != nil {
		http.Error(w, `{"status_code": 500, "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book updated successfully"})
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, `{"status_code": 400, "message": "Invalid ID"}`, http.StatusBadRequest)
		return
	}

	if err := h.bookService.DeleteBook(id); err != nil {
		http.Error(w, `{"status_code": 500, "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
}
