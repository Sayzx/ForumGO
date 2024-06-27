package sql

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "forum.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	fmt.Println("Database connection successfully established")

	// Création de la table utilisateurs
	if _, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		);`); err != nil {
		return nil, fmt.Errorf("failed to create users table: %v", err)
	}

	// Création de la table likes
	if _, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS likes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			post_id INTEGER NOT NULL,
			type TEXT CHECK(type IN ('like', 'dislike')) NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);`); err != nil {
		return nil, fmt.Errorf("failed to create likes table: %v", err)
	}

	return db, nil
}

// Fonction pour ajouter un like ou un dislike
func AddLikeOrDislike(db *sql.DB, userID, postID int, likeType string) error {
	_, err := db.Exec("INSERT INTO likes (user_id, post_id, type) VALUES (?, ?, ?)", userID, postID, likeType)
	return err
}

// Fonction pour compter les likes ou dislikes
func CountLikesOrDislikes(db *sql.DB, postID int, likeType string) (int, error) {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = ? AND type = ?", postID, likeType)
	err := row.Scan(&count)
	return count, err
}
