package handler

import (
	"main/internal/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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
		return
	}

	db, err := sql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (username, email, password, avatar, platform, rank) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, email, hashedPassword, "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png?ex=668913a1&is=6687c221&hm=895af00c0facede320bc213425295dbeae26a1652ae0a217e40a8e80bb418dfe&=&format=webp&quality=lossless&width=640&height=640", "Local", "user")
	if err != nil {
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
