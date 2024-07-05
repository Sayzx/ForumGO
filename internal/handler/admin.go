package handler

import (
	"html/template"
	"log"
	"main/internal/api"
	"main/internal/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type AdminData struct {
	LoggedIn      bool
	ActiveUsers   string
	Avatar        string
	User          User
	Topics        []api.Topic
	ReportedPosts []api.ReportedPost
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	api.GetUsernameByCookie(r)

	// Get the number of active users
	activeUsers := api.GetActiveUsers()
	data := AdminData{
		LoggedIn:    true,
		ActiveUsers: strconv.Itoa(len(activeUsers)),
	}

	cookie, err := r.Cookie("user")
	if err == nil && cookie != nil {
		value, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			log.Println("Error unescaping cookie value:", err)
			http.Error(w, "Error processing cookie", http.StatusBadRequest)
			return
		}

		parts := strings.SplitN(value, ";", 3)
		if len(parts) == 3 {
			data.LoggedIn = true
			data.Avatar = utils.CleanAvatarURL(parts[1])
			data.User = User{Avatar: data.Avatar}
		}
	}

	if !data.LoggedIn {
		data.Avatar = "./web/assets/img/default-avatar.webp"
		data.User = User{Avatar: data.Avatar}
	}

	reportedPosts, err := api.GetReportedPosts()
	if err != nil {
		log.Println("Error getting reported posts:", err)
		http.Error(w, "Error getting reported posts", http.StatusInternalServerError)
		return
	}
	data.ReportedPosts = reportedPosts

	tmpl, err := template.ParseFiles("web/templates/admin.html")
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

func AcceptPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Error parsing post ID:", err)
		http.Error(w, "Error parsing post ID", http.StatusBadRequest)
		return
	}

	err = api.AcceptPost(id)
	if err != nil {
		log.Println("Error accepting post:", err)
		http.Error(w, "Error accepting post", http.StatusInternalServerError)
		return
	}

	// Redirect to the admin page
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
