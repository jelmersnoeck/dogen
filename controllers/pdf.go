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

func PdfTemplate(w http.ResponseWriter, r *http.Request) {
	data := userInput(r)

	template_information, _ := utils.LoadTemplate("farewill")
	template, _ := pdf.NewJsonTemplate(template_information)
	template.LoadBlocks(data)

	f := pdf.NewGoFpdf(template.Layout())
	pdf.ParseBlocks(f, template.Blocks())

	buffer := bytes.NewBufferString("")
	w.Write(f.Bytes(buffer))
}

func userInput(request *http.Request) map[string]interface{} {
	var data map[string]interface{}
	body, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(body, &data)

	if data["data"] == nil {
		return make(map[string]interface{})
	}

	return data["data"].(map[string]interface{})
}

func templateName(request *http.Request) string {
	vars := mux.Vars(request)
	return vars["template"]
}
