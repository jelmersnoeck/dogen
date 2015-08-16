package controllers

import (
	"github.com/jelmersnoeck/noscito/services/easypdf"
	"github.com/jelmersnoeck/noscito/services/easypdf/template"
	"net/http"
)

func MainIndex(w http.ResponseWriter, r *http.Request) {
	template := template.Load("print-batch-collection")

	pdf := easypdf.New(&template.Layout)

	w.Write(easypdf.Render(pdf))
}
