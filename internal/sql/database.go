package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// ConnectDB establishes a connection to the SQLite database
func ConnectDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "forumgo.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %v", err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	fmt.Println("Database connection successfully established")
	return db, nil
}
