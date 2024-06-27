package handler

import (
	"fmt"
	"main/internal/api"
	"net/http"
)

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve the user cookie
	cookie, err := r.Cookie("user")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := api.GetUsernameByCookie(r)
	rank := api.GetGroupByUsername(username)
	fmt.Println(rank)
	fmt.Println(username)
	if rank != "admin" {
		http.Error(w, "You are not an admin", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, "./web/templates/logs.html")
}
