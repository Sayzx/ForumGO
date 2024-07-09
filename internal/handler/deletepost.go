package handler

import (
	"main/internal/api"
	"net/http"
	"strconv"
)

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed ( Genre ta cru tu pouvais delete un post comme sa )", http.StatusMethodNotAllowed)
		return
	}
	// Récupération de l'ID du post à supprimer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Error parsing post ID", http.StatusBadRequest)
		return
	}

	// Tentative de suppression
	err = api.DeletePost(id)
	if err != nil {
		http.Error(w, "Error deleting post", http.StatusInternalServerError)
		return
	}

	// Redirection vers la page d'accueil
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeletePostFromAdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed ( Genre ta cru tu pouvais delete un post comme sa )", http.StatusMethodNotAllowed)
		return
	}
	// Récupération de l'ID du post à supprimer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Error parsing post ID", http.StatusBadRequest)
		return
	}

	// Tentative de suppression
	err = api.DeletePostfromAdmin(id)
	if err != nil {
		http.Error(w, "Error deleting post", http.StatusInternalServerError)
		return
	}

	// Redirection vers la page admin
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
