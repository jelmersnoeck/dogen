package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jelmersnoeck/noscito/controllers"
)

func HandleRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.MainIndex)
	router.HandleFunc("/render/{template}", controllers.PdfTemplate)

	http.Handle("/", router)
}
