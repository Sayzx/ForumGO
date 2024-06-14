package handler

import (
	"main/internal/sql"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// vu qu'on a fait une requet post on va recupeer username email & password
	// on va les recuperer grace a r.FormValue
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	sql.ConnectDB()
	sql.InsertUser(username, email, password)

}
