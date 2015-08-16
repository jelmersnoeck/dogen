package controllers

import (
	"net/http"
)

func MainIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
