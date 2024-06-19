package handler

import (
	"log"

	"main/internal/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer les valeurs du formulaire POST
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Ouvrir la connexion à la base de données
	db, err := sql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Could not connect to the database:", err)
		return
	}
	defer db.Close()

	// Préparer la requête d'insertion
	stmt, err := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		log.Println("Could not prepare query:", err)
		return
	}
	defer stmt.Close()

	// Exécuter la requête d'insertion
	_, err = stmt.Exec(username, email, password)
	if err != nil {
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		log.Println("Could not execute query:", err)
		return
	}

	// Répondre à l'utilisateur
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
