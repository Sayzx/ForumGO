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

func InsertUser(username, email, password string) error {
	// Prepare statement for inserting data
	stmt, err := db.Prepare("INSERT INTO users(username, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare the statement: %v", err)
	}
	defer stmt.Close()

	// Execute the prepared statement
	_, err = stmt.Exec(username, email, password)
	if err != nil {
		return fmt.Errorf("failed to execute the statement: %v", err)
	}

	fmt.Println("New user has been added to the database")
	return nil
}
