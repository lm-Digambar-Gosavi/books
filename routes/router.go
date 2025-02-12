package routes

import (
	"books/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(bookHandler *handlers.BookHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/books", func(r chi.Router) {
		r.Post("/create", bookHandler.CreateBook)
		r.Get("/all", bookHandler.GetAllBooks)
		r.Get("/{id}", bookHandler.GetBookByID)
		r.Put("/{id}", bookHandler.UpdateBook)
		r.Delete("/{id}", bookHandler.DeleteBook)
	})
	return r
}
