package controllers

import (
	"encoding/json"
	"github.com/jelmersnoeck/noscito/services/easypdf"
	"github.com/jelmersnoeck/noscito/services/easypdf/template"
	"net/http"
)

func MainIndex(w http.ResponseWriter, r *http.Request) {
	template := template.Load("print-batch-collection")

	pdf := easypdf.New(template.Layout)
	pdf.RegisterBlocks(template.Blocks, userInput())
	w.Write(pdf.Render())
}

func userInput() (data map[string]interface{}) {
	byt := []byte(`{
		"set_1_image_1": { "url": "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/UPC-A-036000291452.png/220px-UPC-A-036000291452.png", "visible": false },
		"set_1_image_2": { "url": "http://petapixel.com/assets/uploads/2013/11/bloomf1.jpeg" },
		"loop": {
			"items": [
			]
		}
	}`)

	json.Unmarshal(byt, &data)
	return data
}
