package handler

import (
	"main/internal/api"
	dbsql "main/internal/sql"
	"net/http"
	"strconv"
)

func GetIfUserLikedPost(postid int, username string) bool {
	db, err := dbsql.ConnectDB()
	if err != nil {
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM likes WHERE postid = ? AND username = ?", postid, username)
	if err != nil {
		return false
	}
	defer rows.Close()

	return rows.Next()
}

func AddLikeToPost(postid string) {
	db, err := dbsql.ConnectDB()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE topics SET like = like + 1 WHERE id = ?", postid)
	if err != nil {
		return
	}
}

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	postid := r.FormValue("id")
	username := r.FormValue("username")

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("INSERT INTO likes (postid, username) VALUES (?, ?)", postid, username)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE topics SET like = like + 1 WHERE id = ?", postid)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/showpost?postid="+postid, http.StatusSeeOther)
}

func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	postid := r.FormValue("id")
	username := r.FormValue("username")

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("INSERT INTO dislike (postid, username) VALUES (?, ?)", postid, username)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE topics SET dislike = dislike + 1 WHERE id = ?", postid)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/showpost?postid="+postid, http.StatusSeeOther)
}

func GetIfUserHaveDisLike(postid int, username string) bool {
	db, err := dbsql.ConnectDB()
	if err != nil {
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM dislike WHERE postid = ? AND username = ?", postid, username)
	if err != nil {
		return false
	}
	defer rows.Close()

	return rows.Next()
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
