package main

import (
	"net/http"
	"os"

	"github.com/jelmersnoeck/dogen/web/router"
)

func main() {
	router.HandleRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	http.ListenAndServe(":"+port, nil)
}
