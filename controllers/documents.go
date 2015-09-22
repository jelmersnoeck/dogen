package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/jelmersnoeck/noscito/utils"
)

func DocumentsShow(w http.ResponseWriter, r *http.Request) {
	data, inputErr := userInput(r)
	if inputErr != nil {
		w.Write([]byte(inputErr.Error()))
		return
	}

	template, tplErr := loadTemplate(templateName(r), data)
	if tplErr != nil {
		w.Write([]byte(tplErr.Error()))
		return
	}

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

	if data["data"] == nil {
		return make(map[string]interface{}), nil
	}

	return data["data"].(map[string]interface{}), nil
}

func templateName(request *http.Request) string {
	vars := mux.Vars(request)
	return vars["template"]
}

func loadTemplate(name string, userInput map[string]interface{}) (pdf.Template, error) {
	template_information, tplLoadErr := utils.LoadTemplate(name)
	if tplLoadErr != nil {
		return nil, tplLoadErr
	}

	template, jsonErr := pdf.NewJsonTemplate(template_information)

	if jsonErr != nil {
		return nil, jsonErr
	}

	template.LoadBlocks(userInput)

	return template, nil
}
