package handler

import (
	"html/template"
	"main/internal/api"
	"main/internal/sql"
	"net/http"
)

type UserP struct {
	ID       string
	Username string
	Email    string
	Password string
	Rank     string
	Platform string
	Avatar   string
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	username := api.GetUsernameByCookie(r)
	if username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// get all information on db
	db, err := sql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT userid, username, email, password, rank, platform, avatar FROM users WHERE username = ?")
	if err != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var user UserP
	err = stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Rank, &user.Platform, &user.Avatar)
	if err != nil {
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		return
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("./web/templates/profile.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}
}
