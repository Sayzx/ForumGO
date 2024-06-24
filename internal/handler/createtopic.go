package handler

import (
	"fmt"
	"html/template"
	"log"
	"main/internal/api"
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
	var data CreateTopicData

	// Tentative de récupération du cookie utilisateur
	cookie, err := r.Cookie("user")
	if err == nil && cookie != nil {
		value, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			log.Println("Error unescaping cookie value:", err)
			http.Error(w, "Error processing cookie", http.StatusBadRequest)
			return
		}

		parts := strings.SplitN(value, ";", 2)
		if len(parts) == 2 {
			data.LoggedIn = true
			data.Avatar = parts[1]
		}
	}

	if !data.LoggedIn {
		// Définir l'avatar par défaut si l'utilisateur n'est pas connecté
		data.Avatar = "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png"
	}

	// Chargement et exécution du template
	tmpl, err := template.ParseFiles("./web/templates/createtopic.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func AddTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// owner = username by cookie
	username := api.GetUsernameByCookie(r)
	avatar := api.GetAvatarByCookie(r)
	owner := username
	fmt.Println(owner)
	title := r.FormValue("title")
	category := r.FormValue("category")
	tags := r.FormValue("tags")
	content := r.FormValue("content")
	images := r.FormValue("images")
	like := 0
	dislike := 0
	if avatar == "" {
		avatar = "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png?ex=667a9321&is=667941a1&hm=733e73400a7e6e85dac74042fc2ce1f50eeb42c7d53d1228d0dde1e45718fc9d&=&format=webp&quality=lossless"
	}
	if title == "" || category == "" || tags == "" || content == "" || owner == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO topics (title, categoryid, tags, content, images, owner, like, dislike, avatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, category, tags, content, images, owner, like, dislike, avatar)
	if err != nil {
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/showtopics?id="+category, http.StatusSeeOther)
}
