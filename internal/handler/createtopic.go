package handler

import (
	"html/template"
	"io"
	"main/internal/api"
	dbsql "main/internal/sql"
	"main/internal/utils"
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
	// Vérifier si l'utilisateur est connecté
	username := api.GetUsernameByCookie(r)
	if username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var data CreateTopicData

	// Tentative de récupération du cookie utilisateur
	cookie, err := r.Cookie("user")
	if err == nil && cookie != nil {
		value, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			http.Error(w, "Error processing cookie", http.StatusBadRequest)
			return
		}

		parts := strings.SplitN(value, ";", 2)
		if len(parts) == 2 {
			data.LoggedIn = true
			data.Avatar = utils.CleanAvatarURL(parts[1])
		}
	}

	if !data.LoggedIn {
		// Définir l'avatar par défaut si l'utilisateur n'est pas connecté
		data.Avatar = "./web/assets/img/default-avatar.webp"
	}

	// Chargement et exécution du template
	tmpl, err := template.ParseFiles("./web/templates/createtopic.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
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
	title := r.FormValue("title")
	category := r.FormValue("category")
	tags := r.FormValue("tags")
	content := r.FormValue("content")
	like := 0
	dislike := 0
	createat := api.GetDateAndTime()
	if avatar == "" {
		avatar = "./web/assets/img/default-avatar.webp"
	} else {
		avatar = utils.CleanAvatarURL(avatar)
	}
	if title == "" || category == "" || tags == "" || content == "" || owner == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Handle image uploads
	err := r.ParseMultipartForm(MaxUploadSize)
	if err != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}

	var imagePaths []string
	files := r.MultipartForm.File["images[]"] // Note the change here to match HTML form input
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error opening file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fileType := fileHeader.Header.Get("Content-Type")
		if !strings.HasPrefix(fileType, "image/") {
			http.Error(w, "Invalid file type", http.StatusBadRequest)
			return
		}

		if fileHeader.Size > MaxUploadSize {
			http.Error(w, "File is too big", http.StatusBadRequest)
			return
		}

		fileName := filepath.Base(fileHeader.Filename)
		filePath := filepath.Join("/uploads/", fileName)
		filePath = strings.ReplaceAll(filePath, "\\", "/")

		dst, err := os.Create("web/" + filePath)
		if err != nil {
			http.Error(w, "Unable to save the file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Unable to save the file", http.StatusInternalServerError)
			return
		}

		imagePaths = append(imagePaths, filePath)
	}

	images := strings.Join(imagePaths, ";")

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO topics (title, categoryid, tags, content, images, owner, like, dislike, avatar, createat) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, category, tags, content, images, owner, like, dislike, avatar, createat)
	if err != nil {
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/showtopics?id="+category, http.StatusSeeOther)
}
