package database

import (
	"database/sql"
	"encoding/json"
	model "myproject/module"
	"net/http"
)

// CreateBook inserts a new book record into the database.
func CreateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book model.Book

	// Decode the JSON request body into a Book struct.
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}

	// Insert the book record into the "books" table.
	query := "INSERT INTO books(title, author, year) VALUES (?, ?, ?)"
	_, err := db.Exec(query, book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}

	// Respond with the created book record.
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(book)
}

// GetBooks retrieves all book records from the database.
func GetBooks(db *sql.DB) ([]model.Book, error) {
	rows, err := db.Query("SELECT title, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Title, &book.Author, &book.Year); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
