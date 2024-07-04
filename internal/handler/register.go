package handler

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"main/internal/sql"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	if sql.UsernameIsExists(username) {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Println("Error hashing password:", err)
		return
	}

	db, err1 := sql.ConnectDB()
	if err1 != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Could not connect to the database:", err1)
		return
	}
	defer func() {
		if err2 := db.Close(); err2 != nil {
			log.Println("Error closing database:", err2)
		}
	}()

	stmt, err3 := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err3 != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		log.Println("Could not prepare query:", err3)
		return
	}
	defer func() {
		if err4 := stmt.Close(); err4 != nil {
			log.Println("Error closing statement:", err4)
		}
	}()

	_, err5 := stmt.Exec(username, email, hashedPassword)
	if err5 != nil {
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		log.Println("Could not execute query:", err5)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
