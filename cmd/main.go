package main

import (
	"fmt"
	"main/internal/routes"
	"main/internal/utils"
)

func main() {
	// Nettoyage des URLs d'avatars dans la base de données
	utils.CleanDatabaseAvatars()
	routes.Run()
<<<<<<< HEAD

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur: %v\n", err)
	}
=======
>>>>>>> Aylan
}
