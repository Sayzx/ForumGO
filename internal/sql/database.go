package sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
<<<<<<< HEAD
	db, err := sql.Open("sqlite", "forum.db")
=======

	db, err := sql.Open("sqlite3", "internal/sql/forum.db?_busy_timeout=15000&_journal_mode=WAL")
>>>>>>> Aylan
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	createUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		avatar TEXT
	);`

	_, err = db.Exec(createUsersTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %v", err)
	}

	return db, nil
}

func UsernameIsExists(username string) bool {
	db, err := ConnectDB()
	if err != nil {
		return false
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var exists bool
	query := "SELECT COUNT(*) > 0 FROM users WHERE username = ?"
	err = db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		log.Println("Error checking username existence:", err)
		return false
	}
	return exists
}
