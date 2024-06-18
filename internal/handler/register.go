package handler

import (
	sql2 "database/sql"
	"fmt"
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

	fmt.Println("Votre password est : ", password)
	cryptPassword, err1 := HashPassword(password)
	if err1 != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Println("Could not hash password:", err1)
		return
	}
	fmt.Println("Votre password crypté est : ", cryptPassword)

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
	stmt, err4 := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err4 != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		log.Println("Could not prepare query:", err4)
		return
	}
	defer func(stmt *sql2.Stmt) {
		err5 := stmt.Close()
		if err5 != nil {
			log.Println("Could not close the statement:", err5)
		}
	}(stmt)

	// Exécuter la requête d'insertion pour la table users
	res, err6 := stmt.Exec(username, email, cryptPassword)
	if err6 != nil {
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		log.Println("Could not execute query:", err6)
		return
	}

	// Obtenir l'ID de l'utilisateur inséré
	userID, err7 := res.LastInsertId()
	if err7 != nil {
		http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
		log.Println("Could not retrieve last insert ID:", err7)
		return
	}

	// Préparer et exécuter la requête d'insertion pour la table mail
	stmt, err7 = db.Prepare("INSERT INTO mail (user_id, email) VALUES (?, ?)")
	if err7 != nil {
		http.Error(w, "Database query preparation error for mail", http.StatusInternalServerError)
		log.Println("Could not prepare mail query:", err7)
		return
	}
	defer func(stmt *sql2.Stmt) {
		err8 := stmt.Close()
		if err8 != nil {
			log.Println("Could not close the mail statement:", err8)
		}
	}(stmt)

	_, err7 = stmt.Exec(userID, email)
	if err7 != nil {
		http.Error(w, "Database query execution error for mail", http.StatusInternalServerError)
		log.Println("Could not execute mail query:", err7)
		return
	}

	// Préparer et exécuter la requête d'insertion pour la table password
	stmt, err7 = db.Prepare("INSERT INTO password (user_id, password) VALUES (?, ?)")
	if err7 != nil {
		http.Error(w, "Database query preparation error for password", http.StatusInternalServerError)
		log.Println("Could not prepare password query:", err7)
		return
	}
	defer func(stmt *sql2.Stmt) {
		err8 := stmt.Close()
		if err8 != nil {
			log.Println("Could not close the password statement:", err8)
		}
	}(stmt)

	_, err7 = stmt.Exec(userID, cryptPassword)
	if err7 != nil {
		http.Error(w, "Database query execution error for password", http.StatusInternalServerError)
		log.Println("Could not execute password query:", err7)
		return
	}

	// Répondre à l'utilisateur
	_, err7 = w.Write([]byte("User successfully registered"))
	if err7 != nil {
		log.Println("Could not write response:", err7)
		return
	}
}
