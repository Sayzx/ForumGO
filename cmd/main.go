package main

import (
	"main/internal/routes"
	"net/http"
)

func main() {
	routes.Run()
	http.ListenAndServe(":8080", nil)
}
