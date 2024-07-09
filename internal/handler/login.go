package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"main/internal/config"
	dbsql "main/internal/sql"
	"main/internal/utils"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

// create function we take in parameter username and platform if existe on users table

func UserAlreadyRegister(username, platform string) bool {
	db, err := dbsql.ConnectDB()
	if err != nil {
		log.Println("Could not connect to the database:", err)
		return false
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? AND platform = ?", username, platform).Scan(&count)
	if err != nil {
		log.Println("Could not execute query:", err)
		return false
	}

	return count > 0
}

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

	cleanAvatar := utils.CleanAvatarURL(userInfo.Picture)
	userUID := uuid.New().String()
	http.SetCookie(w, &http.Cookie{
		Name:    "user",
		Value:   url.QueryEscape(userInfo.Email + ";" + cleanAvatar + ";" + userUID),
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	})

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Could not connect to the database:", err)
		return
	}
	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Println("Could not close the database connection:", err)
		}
	}(db)

	_, err = db.Exec("INSERT INTO loginlogs (username, platform, datetime) VALUES (?, ?, ?)", userInfo.Email, "Google", time.Now())
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		log.Println("Could not execute query:", err)
		return
	}

	if !UserAlreadyRegister(userInfo.Email, "Google") {
		usernamehash, err := bcrypt.GenerateFromPassword([]byte(userInfo.Email), bcrypt.DefaultCost)
		_, err = db.Exec("INSERT INTO users (userid, username, platform, email, avatar, rank, password) VALUES (?, ?, ?, ?, ?, ?, ?)", userUID, userInfo.Email, "Google", userInfo.Email, cleanAvatar, "user", usernamehash)
		if err != nil {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			log.Println("Could not execute query:", err)
			return
		}
	}
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

	cleanAvatar := utils.CleanAvatarURL(userInfo.AvatarURL)
	userUID := uuid.New().String()

	http.SetCookie(w, &http.Cookie{
		Name:    "user",
		Value:   url.QueryEscape(userInfo.Login + ";" + cleanAvatar + ";" + userUID),
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	})

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Could not connect to the database:", err)
		return
	}
	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Println("Could not close the database connection:", err)
		}
	}(db)

	_, err = db.Exec("INSERT INTO loginlogs (username, platform, datetime) VALUES (?, ?, ?)", userInfo.Login, "GitHub", time.Now())
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		log.Println("Could not execute query:", err)
		return
	}

	if !UserAlreadyRegister(userInfo.Login, "GitHub") {
		usernamehash, err := bcrypt.GenerateFromPassword([]byte(userInfo.Login), bcrypt.DefaultCost)
		_, err = db.Exec("INSERT INTO users (userid, username, platform, email, avatar, rank, password) VALUES (?, ?, ?, ?, ?, ?, ?)", userUID, userInfo.Login, "GitHub", userInfo.Login, cleanAvatar, "user", usernamehash)
		if err != nil {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			log.Println("Could not execute query:", err)
			return
		}
	}
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

	cleanAvatar := utils.CleanAvatarURL(userInfo.Picture.Data.URL)
	userUID := uuid.New().String()

	http.SetCookie(w, &http.Cookie{
		Name:    "user",
		Value:   url.QueryEscape(userInfo.Name + ";" + cleanAvatar + ";" + userUID),
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	})

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Could not connect to the database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO loginlogs (username, platform, datetime) VALUES (?, ?, ?)", userInfo.Name, "Facebook", time.Now())
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		log.Println("Could not execute query:", err)
		return
	}
	if !UserAlreadyRegister(userInfo.Name, "Facebook") {
		usernamehash, err := bcrypt.GenerateFromPassword([]byte(userInfo.Name), bcrypt.DefaultCost)
		_, err = db.Exec("INSERT INTO users (userid, username, platform, email, avatar, rank, password) VALUES (?, ?, ?, ?, ?, ?, ?)", userUID, userInfo.Name, "Facebook", userInfo.Name, cleanAvatar, "user", usernamehash)
		if err != nil {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			log.Println("Could not execute query:", err)
			return
		}
	}
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
		avatarURL = "/web/assets/img/default-avatar.webp"
	}

	cleanAvatar := utils.CleanAvatarURL(avatarURL)
	userUID := uuid.New().String()

	http.SetCookie(w, &http.Cookie{
		Name:    "user",
		Value:   url.QueryEscape(userInfo.Username + ";" + cleanAvatar + ";" + userUID),
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	})

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Could not connect to the database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO loginlogs (username, platform, datetime) VALUES (?, ?, ?)", userInfo.Username, "Discord", time.Now())
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		log.Println("Could not execute query:", err)
		return
	}

	// if UserAlreadyRegister est différent de true alors on insert dans la table users
	if !UserAlreadyRegister(userInfo.Username, "Discord") {
		usernamehash, err := bcrypt.GenerateFromPassword([]byte(userInfo.Username), bcrypt.DefaultCost)
		_, err = db.Exec("INSERT INTO users (userid, username, platform, email, avatar, rank, password) VALUES (?, ?, ?, ?, ?, ?, ?)", userUID, userInfo.Username, "Discord", userInfo.Username, cleanAvatar, "user", usernamehash)
		if err != nil {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			log.Println("Could not execute query:", err)
			return
		}
	}

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

	if CheckPasswordHash(password, storedPasswordHash) {
		userUIID := uuid.New().String()
		cookieValue := url.QueryEscape(storedEmail + ";" + userUIID)
		http.SetCookie(w, &http.Cookie{
			Name:    "user",
			Value:   cookieValue,
			Expires: time.Now().Add(1 * time.Hour),
			Path:    "/",
		})

		_, err = db.Exec("UPDATE users SET userid = ? WHERE username = ? OR email = ?", userUIID, storedEmail, storedEmail)
		if err != nil {
			log.Println("Could not execute query:", err)
			return
		}

		_, err = db.Exec("INSERT INTO loginlogs (username, platform, datetime) VALUES (?, ?, ?)", storedEmail, "Local", time.Now())
		if err != nil {
			log.Println("Could not execute query:", err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "user",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
