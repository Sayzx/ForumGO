package handler

import (
	"html/template"
	"log"
	dbsql "main/internal/sql"
	"net/http"
	"net/url"
	"strings"
)

type CreateTopicData struct {
	LoggedIn bool
	Avatar   string
}

func CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
	data := CreateTopicData{
		Avatar: "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png",
	}

	cookie, err := r.Cookie("user")
	if err == nil && cookie != nil {
		value, err1 := url.QueryUnescape(cookie.Value)
		if err1 != nil {
			log.Println("Error unescaping cookie value:", err1)
			http.Error(w, "Error processing cookie", http.StatusBadRequest)
			return
		}

		parts := strings.SplitN(value, ";", 2)
		if len(parts) == 2 {
			data.LoggedIn = true
			data.Avatar = parts[1]
		}
	}

	tmpl, err2 := template.ParseFiles("./web/templates/createtopic.html")
	if err2 != nil {
		log.Println("Error parsing template:", err2)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	if err3 := tmpl.Execute(w, data); err3 != nil {
		log.Println("Error executing template:", err3)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func AddTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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
	defer func() {
		if err1 := db.Close(); err1 != nil {
			log.Println("Error closing database:", err1)
		}
	}()

	stmt, err2 := db.Prepare("INSERT INTO topics (title, categoryid, tags, content, images) VALUES (?, ?, ?, ?, ?)")
	if err2 != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err3 := stmt.Close(); err3 != nil {
			log.Println("Error closing statement:", err3)
		}
	}()

	_, err4 := stmt.Exec(title, category, tags, content, images)
	if err4 != nil {
		http.Error(w, "Error executing query", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
