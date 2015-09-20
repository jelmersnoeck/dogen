package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

var NonExistingTemplateError = errors.New("Template does not exist.")

func LoadTemplate(name string) ([]byte, error) {
	filepath := path.Join("/app/templates/" + name + ".json")
	fmt.Println(filepath)
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, NonExistingTemplateError
	}

	return file, nil
}
