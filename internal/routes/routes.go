package routes

import (
	"main/internal/handler"
	"net/http"
)

func Run() {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handler.HomeHandler)

	// Admin Page
	// http.HandleFunc("/admin", handler.AdminHandler)

	// Login, Signup, Topic, Post, Comment
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/login-form", handler.LoginFormHandler)
	http.HandleFunc("/signup", handler.SignupHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/createtopic", handler.CreateTopicHandler)
	http.HandleFunc("/addtopic", handler.AddTopicHandler)
	http.HandleFunc("/showtopics", handler.ShowTopicsHandler)
	http.HandleFunc("/showpost", handler.ShowPostHandler)
	http.HandleFunc("/upload", handler.UploadHandler)
	http.HandleFunc("/addcomment", handler.AddCommentHandler)
	http.HandleFunc("/deletepost", handler.DeletePostHandler)
	http.HandleFunc("/like", handler.LikePostHandler)
	http.HandleFunc("/dislike", handler.DislikePostHandler)
	http.HandleFunc("/logs", handler.LogsHandler)

	// GitHub Authentication
	http.HandleFunc("/auth/github", handler.HandleGitHubLogin)
	http.HandleFunc("/github/callback", handler.HandleGitHubCallback)

	http.HandleFunc("/auth/google", handler.HandleGoogleLogin)
	http.HandleFunc("/google/callback", handler.HandleGoogleCallback)

	http.HandleFunc("/auth/facebook", handler.HandleFacebookLogin)
	http.HandleFunc("/facebook/callback", handler.HandleFacebookCallback)

	http.HandleFunc("/auth/discord", handler.HandleDiscordLogin)
	http.HandleFunc("/discord/callback", handler.HandleDiscordCallback)

	http.HandleFunc("/logout", handler.LogoutHandler)
}
