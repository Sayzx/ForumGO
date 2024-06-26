package routes

import (
	"fmt"
	"main/internal/handler"
	"net/http"
)

func Run() {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handler.HomeHandler)

	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/login-form", handler.LoginFormHandler)
	http.HandleFunc("/signup", handler.SignupHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/createtopic", handler.CreateTopicHandler)
	http.HandleFunc("/addtopic", handler.AddTopicHandler)
	http.HandleFunc("/showtopics", handler.ShowTopicsHandler)
	// http.HandleFunc("/showposts", handler.ShowPostsHandler)

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
		fmt.Printf("Erreur lors du démarrage du serveur: %v\n", err)
	}
}
