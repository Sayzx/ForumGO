package handler

import (
	"log"
	"main/internal/api"
	"net/http"
)

func BecomeModeratorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "Missing user ID or username", http.StatusBadRequest)
		log.Println("Missing user ID or username")
		return
	}

	api.BecomeModerator(id)
	// Redirect to the admin page
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
