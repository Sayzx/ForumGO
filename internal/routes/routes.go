package routes

import (
	"fmt"
	"main/internal/handler"
	"net/http"
)

func Run() {
	fmt.Println("Initialisation du serveur...")
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})

	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/signup", handler.SignupHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/auth/github", handler.HandleGitHubLogin)
	http.HandleFunc("/github/callback", handler.HandleGitHubCallback)
	http.HandleFunc("/auth/google", handler.HandleGoogleLogin)
	http.HandleFunc("/callback", handler.HandleGoogleCallback)
	http.HandleFunc("/auth/facebook", handler.HandleFacebookLogin)
	http.HandleFunc("/facebook/callback", handler.HandleFacebookCallback)

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur: %v\n", err)
	}
}
