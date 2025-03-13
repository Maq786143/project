package api

import (
	"database/sql"
	"encoding/json"
	"myproject/database"
	"net/http"
)

// CreateBookLogic processes the creation of a book.
func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {

	return database.CreateBook(db, w, r)
}

// GetBooksLogic retrieves all books and writes them as JSON.
func GetBooksLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	books, err := database.GetBooks(db)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(books)
}
