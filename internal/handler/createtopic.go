package handler

import (
	dbsql "main/internal/sql"
	"net/http"
)

func CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/createtopic.html")
}

func AddTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		title := r.FormValue("title")
		category := r.FormValue("category")
		tags := r.FormValue("tags")
		content := r.FormValue("content")
		images := r.FormValue("images")
		if title == "" || category == "" || tags == "" || content == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}
		if images == "" {
			images = "NULL"
		}

		db, err := dbsql.ConnectDB()
		if err != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		stmt, err := db.Prepare("INSERT INTO topics (title, categoryid, tags, content, images) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			http.Error(w, "Database query preparation error", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
