package handler

import (
	"html/template"
	"log"
	"main/internal/api"
	"main/internal/utils"
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

	topics := api.GetAllTopicsById(categoryid)

	var data struct {
		LoggedIn  bool
		Avatar    string
		Topics    []api.Topic
		Like      int
		Dislike   int
		Username  string
		Createdat string
	}
	data.Topics = topics
	data.Username = username

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
			data.Avatar = utils.CleanAvatarURL(parts[1])
		}
	}

	if !data.LoggedIn {
		data.Avatar = "./web/assets/img/default-avatar.webp"
	}

	// Nettoyer les URLs des avatars dans les topics
	for i := range data.Topics {
		if data.Topics[i].Avatar.Valid {
			data.Topics[i].Avatar.String = utils.CleanAvatarURL(data.Topics[i].Avatar.String)
		}
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
