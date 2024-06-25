package api

import (
	"database/sql"
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
	Avatar       sql.NullString
	CheckLike    sql.NullInt64
	Like         int
	CheckDislike sql.NullInt64
	Dislike      int
	ContentShort string
	CreatedAt    sql.NullString
	Username     string
}

func GetUsernameByCookie(r *http.Request) string {
	cookie, err := r.Cookie("user")
	if err != nil || cookie == nil {
		return ""
	}

	value, err1 := url.QueryUnescape(cookie.Value)
	if err1 != nil {
		return ""
	}

	parts := strings.Split(value, ";")
	return parts[0]
}

func GetAllTopics() []Topic {
	db, err := dbsql.ConnectDB()
	if err != nil {
		log.Println("Could not connect to the database:", err)
		return nil
	}
	defer func() {
		if err1 := db.Close(); err1 != nil {
			log.Println("Could not close the database connection:", err1)
		}
	}()

	stmt, err1 := db.Prepare("SELECT id, title, content, owner, avatar, createat FROM topics")
	if err1 != nil {
		log.Println("Could not prepare query:", err1)
		return nil
	}
	defer func() {
		if err2 := stmt.Close(); err2 != nil {
			log.Println("Could not close the statement:", err2)
		}
	}()

	rows, err2 := stmt.Query()
	if err2 != nil {
		log.Println("Could not execute query:", err2)
		return nil
	}
	defer func() {
		if err3 := rows.Close(); err3 != nil {
			log.Println("Could not close the rows:", err3)
		}
	}()

	var topics []Topic
	for rows.Next() {
		var topic Topic
		err3 := rows.Scan(&topic.ID, &topic.Title, &topic.Content, &topic.Owner, &topic.Avatar, &topic.CreatedAt)
		if err3 != nil {
			log.Println("Could not scan row:", err3)
			return nil
		}
		topics = append(topics, topic)
	}

	if err4 := rows.Err(); err4 != nil {
		log.Println("Error encountered during row iteration:", err4)
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
		if err1 := db.Close(); err1 != nil {
			log.Println("Could not close the database connection:", err1)
		}
	}()

	stmt, err1 := db.Prepare("SELECT id, title, content, owner, avatar, like, dislike FROM topics WHERE categoryid = ?")
	if err1 != nil {
		log.Println("Could not prepare query:", err1)
		return nil
	}
	defer func() {
		if err2 := stmt.Close(); err2 != nil {
			log.Println("Could not close the statement:", err2)
		}
	}()

	rows, err2 := stmt.Query(id)
	if err2 != nil {
		log.Println("Could not execute query:", err2)
		return nil
	}
	defer func() {
		if err3 := rows.Close(); err3 != nil {
			log.Println("Could not close the rows:", err3)
		}
	}()

	var topics []Topic
	for rows.Next() {
		var topic Topic
		err3 := rows.Scan(&topic.ID, &topic.Title, &topic.Content, &topic.Owner, &topic.Avatar, &topic.CheckLike, &topic.CheckDislike)
		if err3 != nil {
			log.Println("Could not scan row:", err3)
			return nil
		}
		topic.Like = int(topic.CheckLike.Int64)
		topic.Dislike = int(topic.CheckDislike.Int64)

		if len(topic.Content) > 50 {
			topic.ContentShort = topic.Content[:50] + "..."
		} else {
			topic.ContentShort = topic.Content
		}
		topics = append(topics, topic)
	}

	if err4 := rows.Err(); err4 != nil {
		log.Println("Error encountered during row iteration:", err4)
		return nil
	}

	return topics
}
