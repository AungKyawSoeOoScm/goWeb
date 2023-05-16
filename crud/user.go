package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	godotenv "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "host=localhost user=postgres password=postakso786 dbname=myuser port=5432 sslmode=disable"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	err := DB.Find(&users).Error
	fmt.Print(err)
	godotenv.Load(".env")
	fmt.Println(os.Getenv("NAME"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// encoder takes a value (usually a struct or a map) and
	// encodes it as JSON, writing the result to the output stream.
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Println(err)
		// Send an HTTP error response with status code 500 (Internal Server Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	err := DB.First(&user, params["id"]).Error
	fmt.Print("err", err)
	if err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	// Decode method to read JSON data from
	// the input stream and decode it into a Go value (e.g., a struct or a map).
	json.NewDecoder(r.Body).Decode(&user)
	err := DB.Create(&user).Error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	err := DB.First(&user, params["id"]).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	err := DB.First(&user, params["id"]).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The user is deleted successfully")
}
