package handler

import (
	"fmt"
	"main/internal/api"
	dbsql "main/internal/sql"
	"net/http"
	"strconv"
)

func GetIfUserLikedPost(postid int, username string) bool {
	db, err := dbsql.ConnectDB()
	if err != nil {
		fmt.Println("Erreur de connexion à la base de données :", err)
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM likes WHERE postid = ? AND username = ?", postid, username)
	if err != nil {
		fmt.Println("Erreur de requête de sélection :", err)
		return false
	}
	defer rows.Close()

	return rows.Next()
}

func AddLikeToPost(postid string) {
	db, err := dbsql.ConnectDB()
	if err != nil {
		fmt.Println("Erreur de connexion à la base de données :", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE topics SET like = like + 1 WHERE id = ?", postid)
	if err != nil {
		fmt.Println("Erreur d'exécution de la requête :", err)
		return
	}
}

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	postid := r.FormValue("id")
	username := r.FormValue("username")
	fmt.Println(postid)
	fmt.Println(username)

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		fmt.Println("Erreur de connexion à la base de données :", err)
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		fmt.Println("Erreur de début de transaction :", err)
		return
	}

	_, err = tx.Exec("INSERT INTO likes (postid, username) VALUES (?, ?)", postid, username)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Database query error", http.StatusInternalServerError)
		fmt.Println("Erreur d'exécution de la requête :", err)
		return
	}

	_, err = tx.Exec("UPDATE topics SET like = like + 1 WHERE id = ?", postid)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Database query error", http.StatusInternalServerError)
		fmt.Println("Erreur d'exécution de la requête de mise à jour :", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		fmt.Println("Erreur de commit de la transaction :", err)
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
		fmt.Println("Erreur de connexion à la base de données :", err)
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		fmt.Println("Erreur de début de transaction :", err)
		return
	}

	_, err = tx.Exec("INSERT INTO dislike (postid, username) VALUES (?, ?)", postid, username)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Database query error", http.StatusInternalServerError)
		fmt.Println("Erreur d'exécution de la requête :", err)
		return
	}

	_, err = tx.Exec("UPDATE topics SET dislike = dislike + 1 WHERE id = ?", postid)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Database query error", http.StatusInternalServerError)
		fmt.Println("Erreur d'exécution de la requête de mise à jour :", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		fmt.Println("Erreur de commit de la transaction :", err)
		return
	}

	http.Redirect(w, r, "/showpost?postid="+postid, http.StatusSeeOther)
}

func GetIfUserHaveDisLike(postid int, username string) bool {
	db, err := dbsql.ConnectDB()
	if err != nil {
		fmt.Println("Erreur de connexion à la base de données :", err)
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM dislike WHERE postid = ? AND username = ?", postid, username)
	if err != nil {
		fmt.Println("Erreur de requête de sélection :", err)
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
		fmt.Println("ID de post invalide :", err)
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
		fmt.Println("Erreur de connexion à la base de données :", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO comments (postid, content, owner, createat, avatar) VALUES (?, ?, ?, ?, ?)", postID, content, username, createat, avatar)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		fmt.Println("Erreur d'exécution de la requête :", err)
		return
	}

	http.Redirect(w, r, "/showpost?postid="+postIDStr, http.StatusSeeOther)
}
