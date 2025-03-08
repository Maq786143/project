package api

import (
	"database/sql"
	"net/http"
)

// RegisterRoutes sets up the HTTP routes.
func RegisterRoutes(db *sql.DB) {

	http.HandleFunc("/create", CreateHandler(db))

	http.HandleFunc("/books", GetBooksHandler(db))
}
