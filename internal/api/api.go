package api

import (
	"log"
	dbsql "main/internal/sql"
	"net/http"
	"net/url"
	"strings"
)

type Author struct {
	Name   string
	Avatar string
}

type Topic struct {
	ID           int
	Title        string
	Content      string
	Owner        string
	Avatar       string
	Like         int
	Dislike      int
	ContentShort string
	CreateAt     *string // Utilisation d'un pointeur pour gérer les valeurs NULL
	Username     string
}

func GetUsernameByCookie(r *http.Request) string {
	cookie, _ := r.Cookie("user")

	value, _ := url.QueryUnescape(cookie.Value)

	parts := strings.Split(value, ";")

	username := parts[0]
	return username
}

func GetAllTopics() []Topic {
	db, err := dbsql.ConnectDB()
	if err != nil {
		log.Println("Could not connect to the database:", err)
		return nil
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Could not close the database connection:", err)
		}
	}()

	stmt, err := db.Prepare("SELECT id, title, content, owner, avatar, createat FROM topics")
	if err != nil {
		log.Println("Could not prepare query:", err)
		return nil
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Println("Could not close the statement:", err)
		}
	}()

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

	var topics []Topic
	for rows.Next() {
		var topic Topic
		err := rows.Scan(&topic.ID, &topic.Title, &topic.Content, &topic.Owner, &topic.Avatar, &topic.CreateAt)
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

func GetAllTopicsById(id string) []Topic {
	db, err := dbsql.ConnectDB()
	if err != nil {
		log.Println("Could not connect to the database:", err)
		return nil
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Could not close the database connection:", err)
		}
	}()

	stmt, err := db.Prepare("SELECT id, title, content, owner, avatar, like, dislike FROM topics where categoryid = ?")
	if err != nil {
		log.Println("Could not prepare query:", err)
		return nil
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Println("Could not close the statement:", err)
		}
	}()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Could not execute query:", err)
		return nil
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Could not close the rows:", err)
		}
	}()

	var topics []Topic
	for rows.Next() {
		var topic Topic
		err := rows.Scan(&topic.ID, &topic.Title, &topic.Content, &topic.Owner, &topic.Avatar, &topic.Like, &topic.Dislike)
		if err != nil {
			log.Println("Could not scan row:", err)
			return nil
		}
		if len(topic.Content) > 50 {
			topic.ContentShort = topic.Content[:50] + "..."
		} else {
			topic.ContentShort = topic.Content
		}
		topics = append(topics, topic)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error encountered during row iteration:", err)
		return nil
	}

	return topics
}

func GetAvatarByCookie(r *http.Request) string {
	cookie, _ := r.Cookie("user")

	value, _ := url.QueryUnescape(cookie.Value)

	parts := strings.Split(value, ";")

	avatar := parts[1]
	return avatar
}
