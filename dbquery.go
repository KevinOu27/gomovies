package main

import (
	"database/sql"
	"fmt"
	"log"
)

// queryMovies fetches movies from the database where the year is greater than 2000.
func queryMovies(db *sql.DB) {
	// Define the SQL query
	query := `SELECT title, year FROM movies WHERE year > 2000;`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close() // Ensure rows cursor is closed after function return

	// Iterate through the result set
	var title string
	var year int
	for rows.Next() {
		err := rows.Scan(&title, &year)
		if err != nil {
			log.Fatalf("Failed to read row: %v", err)
		}
		fmt.Printf("Title: %s, Year: %d\n", title, year)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		log.Fatalf("Error during rows iteration: %v", err)
	}
}
