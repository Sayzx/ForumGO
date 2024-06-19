package handler

import (
	"html/template"
	"log"
	"main/internal/api"
	"net/http"
)

type Author struct {
	Name   string
	Avatar string
}

type Topic struct {
	ID      int
	Title   string
	Content string
	Owner   string
	Avatar  string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	topics := api.GetAllTopics()
	if topics == nil {
		http.Error(w, "Could not fetch topics", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Topics []api.Topic
	}{
		Topics: topics[:min(3, len(topics))], // Display the first three topics
	}

	err = tmpl.Execute(w, data)
	if err != nil {
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
