package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/jelmersnoeck/noscito/utils"
)

func MainIndex(w http.ResponseWriter, r *http.Request) {
	//var data map[string]interface{}
	//body, _ := ioutil.ReadAll(r.Body)
	//json.Unmarshal(body, &data)

	template, _ := pdf.NewJsonTemplate(utils.LoadTemplate("pb-collection"))
	template.LoadBlocks(userInput())
	//template.LoadBlocks(data["data"].(map[string]interface{}))

	f := pdf.NewGoFpdf(template.Layout())
	pdf.ParseBlocks(f, template.Blocks())

	buffer := bytes.NewBufferString("")
	w.Write(f.Bytes(buffer))
}

func userInput() (data map[string]interface{}) {
	byt := []byte(`{
		"set_1_image_1": { "url": "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/UPC-A-036000291452.png/220px-UPC-A-036000291452.png", "visible": false },
		"set_1_image_2": { "url": "http://petapixel.com/assets/uploads/2013/11/bloomf1.jpeg" }
	}`)

	json.Unmarshal(byt, &data)
	return data
}
