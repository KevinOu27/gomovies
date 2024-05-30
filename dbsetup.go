package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func createDatabase() *sql.DB {
	db, err := sql.Open("sqlite", "./gomoviesdb.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS actors (
		id INTEGER PRIMARY KEY,
		name TEXT
	);
	CREATE TABLE IF NOT EXISTS directors (
		id INTEGER PRIMARY KEY,
		name TEXT
	);
	CREATE TABLE IF NOT EXISTS movies (
		id INTEGER PRIMARY KEY,
		title TEXT,
		year INTEGER
	);
	CREATE TABLE IF NOT EXISTS movies_genres (
		movie_id INTEGER,
		genre TEXT,
		FOREIGN KEY(movie_id) REFERENCES movies(id)
	);
	CREATE TABLE IF NOT EXISTS roles (
		id INTEGER PRIMARY KEY,
		movie_id INTEGER,
		actor_id INTEGER,
		role TEXT,
		FOREIGN KEY(movie_id) REFERENCES movies(id),
		FOREIGN KEY(actor_id) REFERENCES actors(id)
	);
	CREATE TABLE IF NOT EXISTS directors_genres (
		director_id INTEGER,
		genre TEXT,
		FOREIGN KEY(director_id) REFERENCES directors(id)
	);
	`
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("Error creating schema: %v", err)
	}

	return db
}
