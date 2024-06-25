package handler

import (
	"html/template"
	"log"
	"main/internal/api"
	"math"
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
	topics := api.GetAllTopics()
	if topics == nil {
		http.Error(w, "Could not fetch topics", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Topics: topics[:int(math.Min(3, float64(len(topics))))],
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

	if !data.LoggedIn {
		data.Avatar = "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png?ex=667a9321&is=667941a1&hm=733e73400a7e6e85dac74042fc2ce1f50eeb42c7d53d1228d0dde1e45718fc9d&=&format=webp&quality=lossless&width=640&height=640"
	}

	tmpl, err2 := template.ParseFiles("./web/templates/index.html")
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
