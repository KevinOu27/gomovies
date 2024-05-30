# MSDS 431 Assignment 10: Personal Movie Database using Go

This repository contains a Go application for managing a movie database. It imports data from CSV files into a SQLite database and allows for querying this data. The application is designed to handle large datasets and enforce database constraints without data duplication.

## System Requirements

- Go (version 1.15 or newer)
- SQLite3
- Access to command-line/terminal

## Setup and Installation

1. **Clone the repository:**
   ```bash
   
   git clone https://github.com/KevinOu27/gomovies.git
   
   cd gomovies
   
   
   
   Then, to build the project: run 'go build .'
   
   To run the application: './gomovies'
   
   (Ensure that all csv files are in the IMDB directory)
   
   Then, run 'go run main.go' to import data and execute predefined queries
   
## Project Structure

dbsetup.go: Sets up the SQLite database and creates the necessary tables.

dataimport.go: Contains functions to import data from CSV files into the database. It handles potential duplicate entries and ensures data integrity.

dbquery.go: Provides functionality to query the database and retrieve information.

main.go: Entry point of the application that orchestrates the creation of the database, data import, and querying operations.

main_test.go: Testing for main.go
  
## Testing 

This application contains unit tests to ensure that database operations are running smoothly.
To run testing, 'go test -v'


## Database Setup and Personalization

### Initial Setup
The SQLite database was set up to efficiently store and query movie data imported from CSV files. This involved creating tables with appropriate data types and constraints to ensure data integrity and performance. The schema was designed to handle various aspects of movie data, including actors, directors, movies, genres, and roles, making it comprehensive for querying diverse movie-related information.

### Personal Movie Collection (Potential Personalizations)
To personalize the application, a new table could be added to the database schema to track individual movie collections. This table would include columns for:
- **Movie ID** (linking to the movies table)

- **Location** (physical location or digital access path)

- **Personal Rating** (user's rating of the movie)

- **Watched Status** (whether the movie has been watched)

- **Notes** (any personal notes or comments on the movie)

Example SQL for creating this table:
```sql
CREATE TABLE personal_collection (
    movie_id INTEGER,
    location TEXT,
    personal_rating INTEGER,
    watched BOOLEAN,
    notes TEXT,
    FOREIGN KEY(movie_id) REFERENCES movies(id)
); ```

### Application Use Cases 

The purpose of integrating the databse is to provide users with a tailored experience to manage, track, and analyze their movie collection separate from the generic data on IMDb. This would allow them to:

- Track which movies they own or have access to

- Note which movies they have watched and thoughts on each one

- Maintain a wishlist of movies to watch or obtain

- Receive recommendations based on preferences and viewing history

### User Interactions

Beyond simple SQL queries, the app could offer an interactive front-end or a command-line interface where users can:

- Add new movies to personal collection

- Update ratings and watched status

- Search for movies based on specific criterias like genre, director, or personal ratings


## Future Enhancements and Developments

## Recommendation System 
A potential enhancement for this application could be the development of a personal movie recommendation system. By analyzing the user's watch history and ratings, we could apply machine learning alogirthms to suggest movies the user may like based on:

- Similarity to highly rated movies by the user 

- Trends in movie genres or directors favored by the user

- Collaborative filtering techniques using data from other users with similar tastes

## Cloud Integration
For accessibility and scalability, the app could be integrated into a cloud platform that would allow users to acces their movie database from any device and share their collections with friends or family. 

## Mobile Application
Developing a mobile app version of the system could provide users with the convenience of managing their collection and receiving recommendations on the go. 

## Conclusion
This personalized movie databse application extends beyond the capabilities of IMDb by offering tailored features to cater towards indiviuals perferences and collection management needs. With further development, it could significantly enhance users' movie-watching experiences and collection management. 
