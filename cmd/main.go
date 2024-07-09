package main

import (
	"main/internal/routes"
	"main/internal/utils"
)

func main() {
	// Nettoyage des URLs d'avatars dans la base de données
	utils.CleanDatabaseAvatars()
	routes.Run()
}
