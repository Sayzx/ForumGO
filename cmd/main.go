package main

import (
	"fmt"
	"main/internal/routes"
	"net/http"
)

func main() {
	routes.Run()

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur: %v\n", err)
	}
}
