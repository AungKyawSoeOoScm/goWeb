package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)

	type Movie struct {
		id          int
		name        string
		genre       string
		releaseYear int
	}
	db, err := sql.Open("mysql", "root:postakso786@(127.0.0.1:3306)/movie?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// query := `
	// CREATE TABLE movies(
	// 	id INT AUTO_INCREMENT,
	// 	name TEXT NOT NULL,
	// 	genre TEXT NOT NULL,
	// 	releaseYear INT NOT NULL,
	// 	PRIMARY KEY (id)
	// )
	// `
	// res, err := db.Exec(query)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// lastId, err := res.LastInsertId()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("The last inserted row id: %d\n", lastId)

	// Insert User
	// movie1 := `
	// INSERT INTO movies(name,genre,releaseYear) VALUES (?,?,?)
	// `
	// result, err := db.Exec(movie1, "Car 2", "Adventure", 2015)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// lastMovieId, err := result.LastInsertId()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("The last inserted row id: %d\n", lastMovieId)

	// Select Specific Movie
	var movie Movie
	query := `SELECT * from movies WHERE id=?`
	serr := db.QueryRow(query, 2).Scan(&movie.id, &movie.name, &movie.genre, &movie.releaseYear)
	if serr != nil {
		log.Fatal(serr)
	}
	fmt.Println("\nGetting Specific Movie ...............")
	fmt.Printf(`Id : %v , Name : %v , Genre: %v , ReleaseYear: %v`, movie.id, movie.name, movie.genre, movie.releaseYear)

	// Select All Movies
	fmt.Println("\nGetting All Movies ..........................")
	allMovie := `SELECT * from movies`
	rows, allerr := db.Query(allMovie)
	if allerr != nil {
		log.Fatal(allerr)
	}
	// defer rows.Close()
	// var movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.id, &movie.name, &movie.genre, &movie.releaseYear)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(`Id : %v , Name : %v , Genre: %v , ReleaseYear: %v`, movie.id, movie.name, movie.genre, movie.releaseYear)
		fmt.Println()
	}
}
