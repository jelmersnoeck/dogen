package main

import (
	"github.com/jelmersnoeck/noscito/router"
	"net/http"
	"os"
)

func main() {
	router.HandleRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	http.ListenAndServe(":"+port, nil)
}
