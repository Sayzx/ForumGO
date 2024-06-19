package handler

import (
	sql2 "database/sql"
	"log"

	"main/internal/api"
	"main/internal/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	err7 := r.ParseForm()
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	cryptPassword, err1 := HashPassword(password)
	if cryptPassword == "" || err1 != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Println("Could not hash password:", err1)
		return
	}
	if username == "" || email == "" || password == "" {
		http.Error(w, "Missing username, email or password", http.StatusBadRequest)
		return
	}
	if err1 != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Println("Could not hash password:", err1)
		return
	}
	db, err2 := sql.ConnectDB()
	if err2 != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Could not connect to the database:", err2)
		return
	}
	defer func(db *sql2.DB) {
		err3 := db.Close()
		if err3 != nil {
			log.Println("Could not close the database connection:", err3)
		}
	}(db)

	// Préparer la requête d'insertion pour la table users
	_, err4 := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err4 != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		log.Println("Could not prepare query:", err4)
		return
	}
	_, err7 = w.Write([]byte("User successfully registered"))
	if err7 != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		log.Println("Could not write response:", err7)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	api.GetAllTopics()

}
