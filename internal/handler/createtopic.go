package handler

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"main/internal/api"
	dbsql "main/internal/sql"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const MaxUploadSize = 20 * 1024 * 1024 // 20 MB
const UploadPath = "./web/uploads"

type CreateTopicData struct {
	LoggedIn bool
	Avatar   string
}

func CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
	var data CreateTopicData

	// Tentative de récupération du cookie utilisateur
	cookie, err := r.Cookie("user")
	if err == nil && cookie != nil {
		value, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			log.Println("Error unescaping cookie value:", err)
			http.Error(w, "Error processing cookie", http.StatusBadRequest)
			return
		}

		parts := strings.SplitN(value, ";", 2)
		if len(parts) == 2 {
			data.LoggedIn = true
			data.Avatar = parts[1]
		}
	}

	if !data.LoggedIn {
		// Définir l'avatar par défaut si l'utilisateur n'est pas connecté
		data.Avatar = "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png"
	}

	// Chargement et exécution du template
	tmpl, err := template.ParseFiles("./web/templates/createtopic.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func AddTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// owner = username by cookie
	username := api.GetUsernameByCookie(r)
	avatar := api.GetAvatarByCookie(r)
	owner := username
	fmt.Println("Owner:", owner)
	title := r.FormValue("title")
	category := r.FormValue("category")
	tags := r.FormValue("tags")
	content := r.FormValue("content")
	like := 0
	dislike := 0
	createat := api.GetDateAndTime()
	if avatar == "" {
		avatar = "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png"
	}
	if title == "" || category == "" || tags == "" || content == "" || owner == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Handle image uploads
	err := r.ParseMultipartForm(MaxUploadSize)
	if err != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		log.Println("File too big:", err)
		return
	}

	// Debugging: Log form data
	log.Println("Form data:", r.MultipartForm)
	log.Println("Form files:", r.MultipartForm.File)

	var imagePaths []string
	files := r.MultipartForm.File["images[]"] // Note the change here to match HTML form input
	log.Println("Number of files received:", len(files))
	for _, fileHeader := range files {
		log.Println("Processing file:", fileHeader.Filename)
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error opening file", http.StatusInternalServerError)
			log.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		fileType := fileHeader.Header.Get("Content-Type")
		if !strings.HasPrefix(fileType, "image/") {
			http.Error(w, "Invalid file type", http.StatusBadRequest)
			log.Println("Invalid file type:", fileType)
			return
		}

		if fileHeader.Size > MaxUploadSize {
			http.Error(w, "File is too big", http.StatusBadRequest)
			log.Println("File is too big:", fileHeader.Size)
			return
		}

		fileName := filepath.Base(fileHeader.Filename)
		filePath := filepath.Join(UploadPath, fileName)

		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Unable to save the file", http.StatusInternalServerError)
			log.Println("Unable to save the file:", err)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Unable to save the file", http.StatusInternalServerError)
			log.Println("Unable to save the file:", err)
			return
		}

		imagePaths = append(imagePaths, filePath)
		log.Println("File uploaded successfully:", filePath)
	}

	images := strings.Join(imagePaths, ";")
	log.Println("Final image paths:", images)

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO topics (title, categoryid, tags, content, images, owner, like, dislike, avatar, createat) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		log.Println("Database query preparation error:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, category, tags, content, images, owner, like, dislike, avatar, createat)
	if err != nil {
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		log.Println("Database query execution error:", err)
		return
	}

	log.Println("Topic created successfully")
	http.Redirect(w, r, "/showtopics?id="+category, http.StatusSeeOther)
}
