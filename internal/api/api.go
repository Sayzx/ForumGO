package api

import (
	"log"
	dbsql "main/internal/sql"
)

type Author struct {
	Name   string
	Avatar string
}

type Topic struct {
	ID      int
	Title   string
	Content string
	Owner   string
	Avatar  string
}

func GetAllTopics() []Topic {
	// Connect to the SQLite 3 database
	db, err := dbsql.ConnectDB() // Use the renamed import
	if err != nil {
		log.Println("Could not connect to the database:", err)
		return nil
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Could not close the database connection:", err)
		}
	}()

	// Prepare the query to get all topics
	stmt, err := db.Prepare("SELECT id, title, content, owner, avatar FROM topics")
	if err != nil {
		log.Println("Could not prepare query:", err)
		return nil
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Println("Could not close the statement:", err)
		}
	}()

	// Execute the query
	rows, err := stmt.Query()
	if err != nil {
		log.Println("Could not execute query:", err)
		return nil
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Could not close the rows:", err)
		}
	}()

	// Process the result
	var topics []Topic
	for rows.Next() {
		var topic Topic
		err := rows.Scan(&topic.ID, &topic.Title, &topic.Content, &topic.Owner, &topic.Avatar)
		if err != nil {
			log.Println("Could not scan row:", err)
			return nil
		}
		topics = append(topics, topic)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error encountered during row iteration:", err)
		return nil
	}

	return topics
}
