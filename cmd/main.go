package main

import (
	"main/internal/routes"
	"main/internal/utils"
	"net/http"
)

func main() {
	// Nettoyage des URLs d'avatars dans la base de données
	utils.CleanDatabaseAvatars()
	routes.Run()
	http.ListenAndServe(":8080", nil)
}
