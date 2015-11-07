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
	"github.com/jelmersnoeck/dogen/renderer/utils"
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

	template, tplErr := loadTemplate(tplName, data)
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
		"text": "JELMER SNOECK",
	}

	template, _ := loadTemplate("demo", data)

	f := pdf.NewGoFpdf(template.Layout())
	pdf.ParseBlocks(f, template.Blocks())

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

func loadTemplate(name string, userInput map[string]interface{}) (templates.Template, error) {
	template_information, tplLoadErr := utils.LoadTemplate(name)
	if tplLoadErr != nil {
		return nil, tplLoadErr
	}

	template, jsonErr := templates.NewJsonTemplate(template_information)

	if jsonErr != nil {
		return nil, jsonErr
	}

	template.LoadBlocks(userInput)

	return template, nil
}
