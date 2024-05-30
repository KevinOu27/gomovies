package main

import (
	"testing"
)

func TestCreateDatabase(t *testing.T) {
	db := createDatabase()
	if db == nil {
		t.Error("Failed to create database")
	}
	db.Close()
}

func TestImportActors(t *testing.T) {
	db := createDatabase()
	defer db.Close()

	importActors(db, "./IMDB/IMDB-actors.csv")

	rows, err := db.Query("SELECT COUNT(*) FROM actors")
	if err != nil {
		t.Errorf("Failed to query actors table: %v", err)
	}

	var count int
	rows.Next()
	rows.Scan(&count)
	rows.Close()

	if count == 0 {
		t.Error("Actors table is empty after import")
	}
}

func TestQueryMovies(t *testing.T) {
	db := createDatabase()
	defer db.Close()

	// Insert sample data
	db.Exec("INSERT INTO movies (id, title, year) VALUES (1, 'Sample Movie', 2021)")

	// Test query
	rows, err := db.Query("SELECT title, year FROM movies WHERE year > 2000")
	if err != nil {
		t.Errorf("Failed to query movies table: %v", err)
	}
	defer rows.Close()

	var title string
	var year int
	if rows.Next() {
		err = rows.Scan(&title, &year)
		if err != nil {
			t.Errorf("Failed to scan row: %v", err)
		}
		if title != "Sample Movie" || year != 2021 {
			t.Errorf("Unexpected result: got (%v, %v)", title, year)
		}
	} else {
		t.Error("No rows returned")
	}
}
