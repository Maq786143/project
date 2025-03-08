package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"myproject/api"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Replace with your actual MySQL credentials and database name
	dsn := "root:Maqbool@123@tcp(127.0.0.1:3306)/library?parseTime=true"

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// Verify the connection is valid
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	fmt.Println("Connection established successfully!")

	// Register API routes
	api.RegisterRoutes(db)
	fmt.Println("Sent the request to routers")

	// Start the HTTP server on port 8080
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
