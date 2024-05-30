package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func importActors(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		log.Fatalf("Error reading header row: %v", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO actors (id, name) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Skipping invalid ID: %v", record[0])
			continue
		}

		_, err = stmt.Exec(id, record[1])
		if err != nil {
			log.Printf("Error inserting actor: %v", err)
			continue
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
func importDirectors(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		log.Fatalf("Error reading header row: %v", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO directors (id, name) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Skipping invalid ID: %v", record[0])
			continue
		}

		_, err = stmt.Exec(id, record[1])
		if err != nil {
			log.Printf("Error inserting director: %v", err)
			continue
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func importMovies(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		log.Fatalf("Error reading header row: %v", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO movies (id, title, year) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		year, err2 := strconv.Atoi(record[2])
		if err != nil || err2 != nil {
			log.Printf("Skipping invalid movie data: ID %v, Year %v", record[0], record[2])
			continue
		}

		_, err = stmt.Exec(id, record[1], year)
		if err != nil {
			log.Printf("Error inserting movie: %v", err)
			continue
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func importMoviesGenres(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		log.Fatalf("Error reading header row: %v", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO movies_genres (movie_id, genre) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, record := range records {
		movieID, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Skipping invalid movie ID: %v", record[0])
			continue
		}

		_, err = stmt.Exec(movieID, record[1])
		if err != nil {
			log.Printf("Error inserting movie genre: %v", err)
			continue
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func importRoles(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		log.Fatalf("Error reading header row: %v", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO roles (id, movie_id, actor_id, role) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		movieID, err1 := strconv.Atoi(record[1])
		actorID, err2 := strconv.Atoi(record[2])
		if err != nil || err1 != nil || err2 != nil {
			log.Printf("Skipping invalid role data: ID %v, Movie ID %v, Actor ID %v", record[0], record[1], record[2])
			continue
		}

		_, err = stmt.Exec(id, movieID, actorID, record[3])
		if err != nil {
			log.Printf("Error inserting role: %v", err)
			continue
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func importDirectorsGenres(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		log.Fatalf("Error reading header row: %v", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO directors_genres (director_id, genre) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, record := range records {
		directorID, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Skipping invalid director ID: %v", record[0])
			continue
		}

		_, err = stmt.Exec(directorID, record[1])
		if err != nil {
			log.Printf("Error inserting director genre: %v", err)
			continue
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
