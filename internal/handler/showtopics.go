package handler

import (
	"fmt"
	"net/http"
)

func ShowTopicsHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'id dans le lien ( id := r.URL.Query().Get("id") )

	categoryid := r.URL.Query().Get("id")
	fmt.Println(categoryid)
	http.ServeFile(w, r, "web/templates/showtopics.html")
}
