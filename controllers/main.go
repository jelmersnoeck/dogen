package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"runtime"

	"github.com/jelmersnoeck/noscito/pdf"
)

func MainIndex(w http.ResponseWriter, r *http.Request) {
	template, _ := pdf.NewJsonTemplate(loadTemplate("print-batch-collection"))
	template.LoadBlocks(userInput())

	f := pdf.NewGoFpdf(template.Layout())
	f.ParseBlocks(template.Blocks())

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

func loadTemplate(name string) []byte {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return nil
	}

	filepath := path.Join(path.Dir(filename), "../templates/"+name+".json")
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	return file
}
