package handler

import (
	"fmt"
	"main/internal/api"
	dbsql "main/internal/sql"
	"net/http"
	"strconv"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	postid := r.FormValue("id")
	username := r.FormValue("username")
	fmt.Println(postid)
	fmt.Println(username)
}

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.FormValue("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	username := api.GetUsernameByCookie(r)
	createat := api.GetDateAndTime()
	avatar := r.FormValue("avatar")

	if content == "" || username == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO comments (postid, content, owner, createat, avatar) VALUES (?, ?, ?, ?, ?)", postID, content, username, createat, avatar)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/showpost?postid="+postIDStr, http.StatusSeeOther)
}
