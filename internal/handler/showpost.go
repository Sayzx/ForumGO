package handler

import (
	"fmt"
	"html/template"
	"log"
	"main/internal/api"
	dbsql "main/internal/sql"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ShowPostData struct {
	LoggedIn    bool
	Avatar      string
	Username    string
	Post        Post
	Comments    []Comment
	IsModerator bool
}

type Post struct {
	ID              int
	Title           string
	Content         string
	Images          string
	Owner           string
	Like            int
	Dislike         int
	CreateAt        string
	UserHaveLike    bool
	UserHaveDislike bool
}

type Comment struct {
	ID       int
	PostID   int
	Content  string
	Owner    string
	CreateAt string
	Avatar   string
}

func ShowPostHandler(w http.ResponseWriter, r *http.Request) {
	var data ShowPostData
	// Attempt to retrieve the user cookie
	cookie, err := r.Cookie("user")
	if err == nil && cookie != nil {
		value, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			log.Println("Error unescaping cookie value:", err)
			http.Error(w, "Error processing cookie", http.StatusBadRequest)
			return
		}

		parts := strings.SplitN(value, ";", 3)
		if len(parts) == 3 {
			data.LoggedIn = true
			data.Username = parts[0]
			data.Avatar = parts[1]
		}
	}

	if !data.LoggedIn {
		data.Avatar = "./web/assets/img/default-avatar.webp"
	}

	// Get user role
	if data.LoggedIn {
		data.IsModerator = api.GetGroupByUsername(data.Username) == "moderator"
	}

	// Retrieve post ID from URL
	postIDStr := r.URL.Query().Get("postid")
	if postIDStr == "" {
		http.Error(w, "Missing post ID 2", http.StatusBadRequest)
		log.Println("Missing post ID 2")
		return
	}

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		log.Println("Invalid post ID:", postIDStr)
		return
	}

	log.Println("Fetching post with ID:", postID)

	// Fetch post details from database
	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	log.Println("Database connection established")

	post := Post{}
	err = db.QueryRow("SELECT id, title, content, images, owner, like, dislike, createat FROM topics WHERE id = ?", postID).Scan(&post.ID, &post.Title, &post.Content, &post.Images, &post.Owner, &post.Like, &post.Dislike, &post.CreateAt)
	fmt.Println(post)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		log.Println("Post not found with ID:", postID)
		log.Println("Error details:", err)
		return
	}
	log.Println("Post found:", post)

	data.Post = post

	// Fetch comments for the post
	rows, err := db.Query("SELECT id, postid, content, owner, createat, avatar FROM comments WHERE postid = ?", postID)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		log.Println("Database query error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.Owner, &comment.CreateAt, &comment.Avatar); err != nil {
			log.Println("Error scanning comment:", err)
			continue
		}
		data.Comments = append(data.Comments, comment)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error during rows iteration:", err)
		http.Error(w, "Error fetching comments", http.StatusInternalServerError)
		return
	}

	log.Println("Comments found:", len(data.Comments))
	HaveLike := GetIfUserLikedPost(postID, data.Username)
	HaveDisLike := GetIfUserHaveDisLike(postID, data.Username)
	if HaveDisLike {
		data.Post.UserHaveDislike = true
	}
	if HaveLike {
		data.Post.UserHaveLike = true
	}

	// Load and execute the template
	tmpl, err := template.ParseFiles("web/templates/showpost.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
