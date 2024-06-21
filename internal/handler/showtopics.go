package handler

import (
	"html/template"
	"log"
	"main/internal/api"
	"net/http"
	"net/url"
	"strings"
)

func ShowTopicsHandler(w http.ResponseWriter, r *http.Request) {
	username := api.GetUsernameByCookie(r)
	if username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	categoryid := r.URL.Query().Get("id")

	// Récupérer les topics par catégorie
	topics := api.GetAllTopicsById(categoryid)
	if topics == nil {
		http.Error(w, "Could not fetch topics", http.StatusInternalServerError)
		return
	}

	// Préparer les données pour le template
	var data struct {
		LoggedIn bool
		Avatar   string
		Topics   []api.Topic
		Like     int
		Dislike  int
		Username string
	}
	data.Topics = topics
	data.Username = username

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

	// Charger et exécuter le template showtopics.html
	tmpl, err := template.ParseFiles("web/templates/showtopics.html")
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
