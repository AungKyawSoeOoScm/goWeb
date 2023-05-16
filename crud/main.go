package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/polar/", http.StripPrefix("/static/", fs))
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(fs)
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	http.Handle("/", r)

	// Start the server
	http.ListenAndServe(":9000", nil)

}
func main() {
	InitialMigration()
	initializeRouter()
}
