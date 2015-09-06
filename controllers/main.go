package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jelmersnoeck/noscito/pdf"
)

func MainIndex(w http.ResponseWriter, r *http.Request) {
	layout, _ := pdf.NewPageLayout("L", "mm", 50, 50)
	f := pdf.NewGoFpdf(layout)

	image := &pdf.Image{
		"http://4.bp.blogspot.com/-JOqxgp-ZWe0/U3BtyEQlEiI/AAAAAAAAOfg/Doq6Q2MwIKA/s1600/google-logo-874x288.png",
		0,
		0,
		10,
		10,
	}
	blocks := make([]pdf.Block, 1)
	blocks[0] = image

	f.ParseBlocks(blocks)

	buffer := bytes.NewBufferString("")
	w.Write(f.Bytes(buffer))
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
