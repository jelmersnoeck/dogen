package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
)

var NonExistingTemplateError = errors.New("Template does not exist.")

func LoadTemplate(name string) ([]byte, error) {
	filepath := path.Join(os.Getwd() + "/templates/" + name + ".json")
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, NonExistingTemplateError
	}

	return file, nil
}
