package handler

import "net/http"

func ShowTopicsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/showtopics.html")
}
