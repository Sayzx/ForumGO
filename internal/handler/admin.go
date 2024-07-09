package handler

import (
	"html/template"
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
	Users         []api.User
	Rank          string
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	username := api.GetUsernameByCookie(r)
	rank := api.GetGroupByUsername(username)
	Rank := strings.ToLower(rank)
	if Rank == "user" {
		http.Error(w, "You are not an admin", http.StatusForbidden)
		return
	}

	// Get the number of active users
	activeUsers := api.GetActiveUsers()
	data := AdminData{
		LoggedIn:    true,
		ActiveUsers: strconv.Itoa(len(activeUsers)),
		Rank:        rank,
	}

	cookie, err := r.Cookie("user")
	if err == nil && cookie != nil {
		value, err := url.QueryUnescape(cookie.Value)
		if err != nil {
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
		http.Error(w, "Error getting reported posts", http.StatusInternalServerError)
		return
	}
	data.ReportedPosts = reportedPosts

	users, err := api.GetAllUsers()
	if err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}
	data.Users = users

	tmpl, err := template.ParseFiles("web/templates/admin.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
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
		http.Error(w, "Error parsing post ID", http.StatusBadRequest)
		return
	}

	err = api.AcceptPost(id)
	if err != nil {
		http.Error(w, "Error accepting post", http.StatusInternalServerError)
		return
	}

	// Redirect to the admin page
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Error parsing user ID", http.StatusBadRequest)
		return
	}

	err = api.DeleteUser(id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	// Redirect to the admin page
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
