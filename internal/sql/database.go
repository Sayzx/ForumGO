package sql

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "internal/sql/forum.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		if err1 := db.Close(); err1 != nil {
			return nil, err1
		}
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS mail (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			email TEXT NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);`,
		`CREATE TABLE IF NOT EXISTS password (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			password TEXT NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);`,
	}

	for _, query := range queries {
		if _, err = db.Exec(query); err != nil {
			return nil, fmt.Errorf("failed to execute query: %v", err)
		}
	}

	return db, nil
}
