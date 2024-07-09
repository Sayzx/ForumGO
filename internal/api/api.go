package api

import (
	"database/sql"
	"log"
	dbsql "main/internal/sql"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	CreateAt     *string // Utilisation d'un pointeur pour gérer les valeurs NULL
	Username     string
}

type ReportedPost struct {
	ID      int
	Title   string
	Content string
	Owner   string
	Avatar  sql.NullString
}

func GetUsernameByCookie(r *http.Request) string {
	cookie, _ := r.Cookie("user")

	if cookie == nil {
		return ""
	}

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

func GetDateAndTime() string {
	// get today date and time
	now := time.Now()
	return now.Format("02-01-2006 15:04")
}

func GetActiveUsers() []Author {
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

	rows, err := db.Query("SELECT username, avatar FROM users WHERE active = 1")
	if err != nil {
		log.Println("Could not query users:", err)
		return nil
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Could not close the rows:", err)
		}
	}()

	var authors []Author
	for rows.Next() {
		var author Author
		err := rows.Scan(&author.Name, &author.Avatar)
		if err != nil {
			log.Println("Could not scan row:", err)
			return nil
		}
		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error encountered during row iteration:", err)
		return nil
	}

	return authors
}

func DeletePost(id int) error {
	db, err := dbsql.ConnectDB()
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Could not close the database connection:", err)
		}
	}()

	stmt, err := db.Prepare("DELETE FROM topics WHERE id = ?")
	if err != nil {
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Println("Could not close the statement:", err)
		}
	}()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
func DeletePostfromAdmin(id int) error {
	db, err := dbsql.ConnectDB()
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Could not close the database connection:", err)
		}
	}()

	// Delete from reportspost table
	stmt1, err := db.Prepare("DELETE FROM reportspost WHERE postid = ?")
	if err != nil {
		return err
	}
	defer func() {
		if err := stmt1.Close(); err != nil {
			log.Println("Could not close the statement:", err)
		}
	}()

	_, err = stmt1.Exec(id)
	if err != nil {
		return err
	}

	// Delete from topics table
	stmt2, err := db.Prepare("DELETE FROM topics WHERE id = ?")
	if err != nil {
		return err
	}
	defer func() {
		if err := stmt2.Close(); err != nil {
			log.Println("Could not close the statement:", err)
		}
	}()

	_, err = stmt2.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func GetGroupByUsername(username string) string {
	// select rank from users where username = username
	db, err := dbsql.ConnectDB()
	if err != nil {
		log.Println("Could not connect to the database:", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Could not close the database connection:", err)
		}
	}()

	stmt, err := db.Prepare("SELECT rank FROM users WHERE email = ?")
	if err != nil {
		log.Println("Could not prepare query:", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Println("Could not close the statement:", err)
		}
	}()

	rows, err := stmt.Query(username)
	if err != nil {
		log.Println("Could not execute query:", err)
	}

	var rank string
	for rows.Next() {
		err := rows.Scan(&rank)
		if err != nil {
			log.Println("Could not scan row:", err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Println("Error encountered during row iteration:", err)
	}

	return rank
}

func GetReportedPosts() ([]ReportedPost, error) {
	db, err := dbsql.ConnectDB()
	if err != nil {
		log.Println("Could not connect to the database:", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT postid, title, content, owner, avatar FROM reportspost")
	if err != nil {
		log.Println("Could not query reported posts:", err)
		return nil, err
	}
	defer rows.Close()

	var reportedPosts []ReportedPost
	for rows.Next() {
		var post ReportedPost
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Owner, &post.Avatar)
		if err != nil {
			log.Println("Could not scan row:", err)
			return nil, err
		}
		reportedPosts = append(reportedPosts, post)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error encountered during row iteration:", err)
		return nil, err
	}

	return reportedPosts, nil
}

func AcceptPost(id int) error {
	db, err := dbsql.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Supprimer le post de la table "reportspost"
	stmt, err := db.Prepare("DELETE FROM reportspost WHERE postid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
