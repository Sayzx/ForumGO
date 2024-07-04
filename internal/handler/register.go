package handler

import (
	"log"
	"net/http"
	"strings"

	"main/internal/sql"

	"golang.org/x/crypto/bcrypt"
)

// Simulated function to fetch email content; in practice, replace with your actual method.
func fetchEmailContent() string {
	return `From: sender@example.com
To: receiver@example.com
Subject: New User Registration
Username: newuser
Email: newuser@example.com
Password: test1234`
}

func extractEmailData(emailContent string) (username, email, password string) {
	lines := strings.Split(emailContent, "\n")
	data := make(map[string]string)
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			data[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return data["Username"], data["Email"], data["Password"]
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	emailContent := fetchEmailContent()
	username, email, password := extractEmailData(emailContent)

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

	db, err := sql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Could not connect to the database:", err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		log.Println("Could not prepare query:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, email, hashedPassword)
	if err != nil {
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		log.Println("Could not execute query:", err)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
