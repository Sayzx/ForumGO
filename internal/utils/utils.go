package utils

import (
	"log"
	dbsql "main/internal/sql"
	"net/url"
	"strings"
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
	return cleanURL
}

// CleanDatabaseAvatars nettoie les URLs d'avatars dans la base de données
func CleanDatabaseAvatars() {
	db, err := dbsql.ConnectDB()
	if err != nil {
		log.Println("Could not connect to the database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, avatar FROM users WHERE avatar IS NOT NULL")
	if err != nil {
		log.Println("Could not query avatars:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var avatar string
		err := rows.Scan(&id, &avatar)
		if err != nil {
			log.Println("Could not scan row:", err)
			continue
		}

		cleanAvatar := CleanAvatarURL(avatar)
		_, err = db.Exec("UPDATE users SET avatar = ? WHERE id = ?", cleanAvatar, id)
		if err != nil {
			log.Println("Could not update avatar:", err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Println("Error encountered during row iteration:", err)
	}
}
