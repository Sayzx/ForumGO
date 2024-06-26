package handler

import (
	"main/internal/api"
	dbsql "main/internal/sql"
	"net/http"
	"strconv"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.FormValue("postid")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE topics SET like = like + 1 WHERE id = ?", postID)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.FormValue("postid")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE topics SET dislike = dislike + 1 WHERE id = ?", postID)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.FormValue("postid")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	username := api.GetUsernameByCookie(r)
	createat := api.GetDateAndTime()

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

	_, err = db.Exec("INSERT INTO comments (postid, content, owner, createat) VALUES (?, ?, ?, ?)", postID, content, username, createat)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/showpost?id="+postIDStr, http.StatusSeeOther)
}
