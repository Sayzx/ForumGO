package handler

import "net/http"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/login.html")
}
