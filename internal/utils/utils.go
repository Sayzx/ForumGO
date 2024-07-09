package utils

import (
	dbsql "main/internal/sql"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

type OAuthProfile struct {
	AvatarURL string `json:"avatar_url"`
}

// CleanAvatarURL nettoie l'URL de l'avatar en retirant les parties indésirables après le premier "="
func CleanAvatarURL(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		// En cas d'erreur, retourner l'URL brute
		return rawURL
	}
	cleanURL := strings.TrimSpace(parsedURL.String())
	// Ajoutez un paramètre aléatoire pour éviter le cache
	randomParam := GenerateRandomString(10)
	cleanURL = cleanURL + "?rand=" + randomParam
	return cleanURL
}

// CleanDatabaseAvatars nettoie les URLs d'avatars dans la base de données
func CleanDatabaseAvatars() {
	db, err := dbsql.ConnectDB()
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, avatar FROM users WHERE avatar IS NOT NULL")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var avatar string
		rows.Scan(&id, &avatar)

		cleanAvatar := CleanAvatarURL(avatar)
		_, err = db.Exec("UPDATE users SET avatar = ? WHERE id = ?", cleanAvatar, id)
	}

}

// GenerateRandomString génère une chaîne aléatoire de longueur n
func GenerateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
