package handler

import (
	"database/sql"
	"fmt"
	"main/internal/api"
	dbsql "main/internal/sql"
	"net/http"
)

type ReportedPost struct {
	PostID  int
	Content string
	Owner   string
	Title   string
	Avatar  sql.NullString
}

func ReportPostHandler(w http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve the user cookie
	cookie, err := r.Cookie("user")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := api.GetUsernameByCookie(r)
	rank := api.GetGroupByUsername(username)
	if rank == "user" || rank == "" {
		http.Error(w, "You are not an admin", http.StatusForbidden)
		return
	}

	postid := r.URL.Query().Get("id")
	postcontent := r.FormValue("content")
	postowner := r.FormValue("owner")
	posttitle := r.FormValue("title")
	avatar := r.FormValue("avatar")

	// insert on reportspost table on db
	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		fmt.Println("Erreur de connexion à la base de données :", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO reportspost (postid, content, owner, title, avatar) VALUES (?, ?, ?, ?, ?)", postid, postcontent, postowner, posttitle, avatar)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		fmt.Println("Erreur d'exécution de la requête :", err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetReportedPosts() ([]ReportedPost, error) {
	db, err := dbsql.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT postid, content, owner, title, avatar FROM reportspost")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []ReportedPost
	for rows.Next() {
		var post ReportedPost
		if err := rows.Scan(&post.PostID, &post.Content, &post.Owner, &post.Title, &post.Avatar); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
