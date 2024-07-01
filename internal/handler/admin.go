// package handler

// import (
// 	"html/template"
// 	"log"
// 	"main/internal/api"
// 	"main/internal/utils"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"strings"
// )

// type User struct {
// 	Avatar string
// }

// type AdminData struct {
// 	ActiveUsers string
// 	Avatar      string
// 	Topics      []api.Topic
// 	User        User
// }

// func AdminHandler(w http.ResponseWriter, r *http.Request) {
// 	api.GetUsernameByCookie(r)

// 	// Get the number of active users
// 	activeUsers := api.GetActiveUsers()
// 	data := AdminData{
// 		ActiveUsers: strconv.Itoa(len(activeUsers)),
// 	}

// 	cookie, err := r.Cookie("user")
// 	if err == nil && cookie != nil {
// 		value, err := url.QueryUnescape(cookie.Value)
// 		if err != nil {
// 			log.Println("Error unescaping cookie value:", err)
// 			http.Error(w, "Error processing cookie", http.StatusBadRequest)
// 			return
// 		}

// 		log.Println("Cookie value:", value)
// 		parts := strings.SplitN(value, ";", 3)
// 		if len(parts) == 3 {
// 			PageData.LoggedIn = true
// 			data.Avatar = utils.CleanAvatarURL(parts[1])
// 			data.User = User{Avatar: data.Avatar}
// 			log.Println("Avatar URL after cleaning:", data.Avatar)
// 		}
// 	} else {
// 		log.Println("No valid user cookie found, user not logged in.")
// 	}

// 	if !PageData.LoggedIn {
// 		data.Avatar = "./web/assets/img/default-avatar.webp"
// 		data.User = User{Avatar: data.Avatar}
// 	}

// 	tmpl, err := template.ParseFiles("web/templates/admin.html")
// 	if err != nil {
// 		log.Println("Error parsing template:", err)
// 		http.Error(w, "Error parsing template", http.StatusInternalServerError)
// 		return
// 	}

// 	if err := tmpl.Execute(w, data); err != nil {
// 		log.Println("Error executing template:", err)
// 		http.Error(w, "Error executing template", http.StatusInternalServerError)
// 	}
// }
