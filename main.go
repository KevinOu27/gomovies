package main

func main() {
	db := createDatabase()
	defer db.Close()

	// Import data from CSV files
	importActors(db, "./IMDB/IMDB-actors.csv")
	importDirectors(db, "./IMDB/IMDB-directors.csv")
	importMovies(db, "./IMDB/IMDB-movies.csv")
	importMoviesGenres(db, "./IMDB/IMDB-movies_genres.csv")
	importRoles(db, "./IMDB/IMDB-roles.csv")
	importDirectorsGenres(db, "./IMDB/IMDB-directors_genres.csv")

	// Querying the database
	queryMovies(db)
}
