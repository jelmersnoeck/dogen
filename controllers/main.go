package controllers

import (
	"github.com/jelmersnoeck/noscito/services/easypdf"
	"github.com/jelmersnoeck/noscito/services/easypdf/template"
	"net/http"
)

func MainIndex(w http.ResponseWriter, r *http.Request) {
	template := template.Load("print-batch-collection")

	pdf := easypdf.New(&template.Layout)

	for _, block := range template.Blocks {
		block.Item.Parse(pdf)
	}

	w.Write(easypdf.Render(pdf))
}
