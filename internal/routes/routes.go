package routes

import (
	"fmt"
	"main/internal/handler"
	"net/http"
)

func Run() {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fs2 := http.FileServer(http.Dir("web/uploads"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs2))

	http.HandleFunc("/", handler.HomeHandler)

	// Admin Page
	http.HandleFunc("/admin", handler.AdminHandler)

	// Login, Signup, Topic, Post, Comment, Report, Delete, Accept, Like, Dislike, Logs
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
	http.HandleFunc("/reportpost", handler.ReportPostHandler)
	http.HandleFunc("/deletepost", handler.DeletePostHandler)
	http.HandleFunc("/deletepostfromadmin", handler.DeletePostFromAdminHandler)
	http.HandleFunc("/acceptpost", handler.AcceptPostHandler)
	http.HandleFunc("/like", handler.LikePostHandler)
	http.HandleFunc("/dislike", handler.DislikePostHandler)
	http.HandleFunc("/logs", handler.LogsHandler)
	http.HandleFunc("/profile", handler.ProfileHandler)
	http.HandleFunc("/deleteuser", handler.DeleteUserHandler)
	http.HandleFunc("/becomemod", handler.BecomeModeratorHandler)

	// GitHub Authentication
	http.HandleFunc("/auth/github", handler.HandleGitHubLogin)
	http.HandleFunc("/github/callback", handler.HandleGitHubCallback)

	// Google Authentication
	http.HandleFunc("/auth/google", handler.HandleGoogleLogin)
	http.HandleFunc("/google/callback", handler.HandleGoogleCallback)

	// Facebook Authentication
	http.HandleFunc("/auth/facebook", handler.HandleFacebookLogin)
	http.HandleFunc("/facebook/callback", handler.HandleFacebookCallback)

	// Discord Authentication
	http.HandleFunc("/auth/discord", handler.HandleDiscordLogin)
	http.HandleFunc("/discord/callback", handler.HandleDiscordCallback)

	// Logout
	http.HandleFunc("/logout", handler.LogoutHandler)

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
	}
}
