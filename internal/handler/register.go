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

<<<<<<< HEAD
	stmt, err3 := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err3 != nil {
=======
	stmt, err := db.Prepare("INSERT INTO users (username, email, password, avatar, platform, rank) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
>>>>>>> Aylan
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		log.Println("Could not prepare query:", err3)
		return
	}
	defer func() {
		if err4 := stmt.Close(); err4 != nil {
			log.Println("Error closing statement:", err4)
		}
	}()

<<<<<<< HEAD
	_, err5 := stmt.Exec(username, email, hashedPassword)
	if err5 != nil {
=======
	_, err = stmt.Exec(username, email, hashedPassword, "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png?ex=668913a1&is=6687c221&hm=895af00c0facede320bc213425295dbeae26a1652ae0a217e40a8e80bb418dfe&=&format=webp&quality=lossless&width=640&height=640", "Local", "user")
	if err != nil {
>>>>>>> Aylan
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		log.Println("Could not execute query:", err5)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
