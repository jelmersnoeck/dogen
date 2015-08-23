package controllers

import (
	"encoding/json"
	"github.com/jelmersnoeck/noscito/services/easypdf"
	"github.com/jelmersnoeck/noscito/services/easypdf/template"
	"net/http"
)

func MainIndex(w http.ResponseWriter, r *http.Request) {
	template := template.Load("print-batch-collection")

	data := userInput()
	template.MatchData(data)

	pdf := easypdf.New(template.Layout)
	easypdf.LoadBlocks(pdf, template.Blocks, data)

	w.Write(easypdf.Render(pdf))
}

func userInput() (data map[string]interface{}) {
	byt := []byte(`{
		"set_1_image_1": "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/UPC-A-036000291452.png/220px-UPC-A-036000291452.png",
		"set_1_image_2": "http://petapixel.com/assets/uploads/2013/11/bloomf1.jpeg"
	}`)

	json.Unmarshal(byt, &data)
	return data
}
