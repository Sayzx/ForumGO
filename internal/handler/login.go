package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"main/internal/config"
	dbsql "main/internal/sql"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := config.GoogleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := config.GitHubOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	url := config.FacebookOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleDiscordLogin(w http.ResponseWriter, r *http.Request) {
	url := config.DiscordOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := r.FormValue("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}
	token, err := config.GoogleOauthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	userInfo := struct {
		Email   string `json:"email"`
		Picture string `json:"picture"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set a cookie with user info
	http.SetCookie(w, &http.Cookie{
		Name:    "user",
		Value:   url.QueryEscape(userInfo.Email + ";" + userInfo.Picture),
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	})

	// Redirect to home page or send the user info as needed
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := r.FormValue("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}
	token, err := config.GitHubOauthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	client := config.GitHubOauthConfig.Client(ctx, token)
	user, err := client.Get("https://api.github.com/user")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer user.Body.Close()

	userInfo := struct {
		Login     string `json:"login"`
		AvatarURL string `json:"avatar_url"`
	}{}

	if err := json.NewDecoder(user.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "user",
		Value:   url.QueryEscape(userInfo.Login + ";" + userInfo.AvatarURL),
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleFacebookCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := r.FormValue("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}
	token, err := config.FacebookOauthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	client := config.FacebookOauthConfig.Client(ctx, token)
	user, err := client.Get("https://graph.facebook.com/me?fields=id,name,email,picture")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer user.Body.Close()

	userInfo := struct {
		Name    string `json:"name"`
		Picture struct {
			Data struct {
				URL string `json:"url"`
			} `json:"data"`
		} `json:"picture"`
	}{}

	if err := json.NewDecoder(user.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "user",
		Value:   url.QueryEscape(userInfo.Name + ";" + userInfo.Picture.Data.URL),
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleDiscordCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := r.FormValue("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}
	token, err := config.DiscordOauthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	client := config.DiscordOauthConfig.Client(ctx, token)
	user, err := client.Get("https://discord.com/api/users/@me")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer user.Body.Close()

	userInfo := struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		ID       string `json:"id"`
	}{}

	if err := json.NewDecoder(user.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	avatarURL := ""
	if userInfo.Avatar != "" {
		avatarURL = fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", userInfo.ID, userInfo.Avatar)
	} else {
		// Default avatar if none is available
		avatarURL = "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png?ex=6673fba1&is=6672aa21&hm=5741edc76eb55c2e3e4ac8924a89c2d610df57a88caf4880636b97a92b3fc153&format=webp&quality=lossless&width=640&height=640&"
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "user",
		Value:   url.QueryEscape(userInfo.Username + ";" + avatarURL),
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/sign-in.html")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/login.html")
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login form submitted")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		log.Println("Failed to parse form:", err)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println("Email:", email)
	fmt.Println("Password:", password)

	if email == "" || password == "" {
		http.Error(w, "Missing email or password", http.StatusBadRequest)
		return
	}

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Could not connect to the database:", err)
		return
	}
	defer db.Close()

	var storedEmail, storedPasswordHash string
	err = db.QueryRow("SELECT email, password FROM users WHERE email = ?", email).Scan(&storedEmail, &storedPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		} else {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			log.Println("Could not execute query:", err)
		}
		return
	}
	fmt.Println("Stored Email:", storedEmail)
	fmt.Println("Stored Password Hash:", storedPasswordHash)

	if CheckPasswordHash(password, storedPasswordHash) {
		// Successful login
		fmt.Println("Login successful!")
		// Set a default avatar cookie here, as user does not have an avatar URL.
		http.SetCookie(w, &http.Cookie{
			Name:    "user",
			Value:   url.QueryEscape(storedEmail + ";https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png?ex=6673fba1&is=6672aa21&hm=5741edc76eb55c2e3e4ac8924a89c2d610df57a88caf4880636b97a92b3fc153&format=webp&quality=lossless&width=640&height=640&"),
			Expires: time.Now().Add(24 * time.Hour),
			Path:    "/",
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
	}
}
