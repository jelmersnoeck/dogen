package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jelmersnoeck/dogen/renderer"
	"github.com/jelmersnoeck/dogen/renderer/documents/pdf"
	"github.com/jelmersnoeck/dogen/renderer/templates"
)

func DocumentsShow(w http.ResponseWriter, r *http.Request) {
	tplName := templateName(r)
	log.Println("Loading " + tplName)

	data, inputErr := userInput(r)
	if inputErr != nil {
		log.Println(tplName + ": " + inputErr.Error())
		w.Write([]byte(inputErr.Error()))
		return
	}

	template, tplErr := templates.LoadJsonTemplate(tplName, data)
	if tplErr != nil {
		log.Println(tplName + ": " + tplErr.Error())
		w.Write([]byte(tplErr.Error()))
		return
	}

	log.Println("Parsing " + tplName)
	f := pdf.NewGoFpdf(template.Layout())
	renderer.Render(template, f.Document())

	log.Println("Displaying " + tplName)
	buffer := bytes.NewBufferString("")
	w.Write(f.Bytes(buffer))
}

func DemoShow(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"demo_text":  map[string]interface{}{"text": "JELMER SNOECK"},
		"demo_image": map[string]interface{}{"url": "https://golang.org/doc/gopher/frontpage.png"},
	}

	template, _ := templates.LoadJsonTemplate("demo", data)

	f := pdf.NewGoFpdf(template.Layout())
	renderer.Render(template, f.Document())

	buffer := bytes.NewBufferString("")
	w.Write(f.Bytes(buffer))
}

func userInput(request *http.Request) (map[string]interface{}, error) {
	var data map[string]interface{}
	body, readErr := ioutil.ReadAll(request.Body)
	json.Unmarshal(body, &data)

	if readErr != nil {
		return nil, readErr
	}

	return data, nil
}

func templateName(request *http.Request) string {
	vars := mux.Vars(request)
	return vars["template"]
}
