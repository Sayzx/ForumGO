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

type PageData struct {
	LoggedIn bool
	Avatar   string
	Topics   []api.Topic
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Chargement des topics
	topics := api.GetAllTopics()

	var data PageData
	data.Topics = topics[:min(3, len(topics))]

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

	// Chargement et exécution du template
	tmpl, err := template.ParseFiles("./web/templates/index.html")
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
