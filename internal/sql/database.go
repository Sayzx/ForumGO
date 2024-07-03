package sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "internal/sql/forum.db?_busy_timeout=15000&_journal_mode=WAL")
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		err1 := db.Close()
		if err1 != nil {
			return nil, err1
		}
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	fmt.Println("Database connection successfully established")

	createUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
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
	defer db.Close()

	var exists bool
	query := "SELECT COUNT(*) > 0 FROM users WHERE username = ?"
	err = db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		log.Println("Error checking username existence:", err)
		return false
	}
	return exists
}
