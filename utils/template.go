package utils

import (
	"errors"
	"io/ioutil"
	"path"
	"runtime"
)

var NonExistingTemplateError = errors.New("Template does not exist.")

func LoadTemplate(name string) ([]byte, error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return nil, errors.New("Could not find current path.")
	}

	filepath := path.Join(path.Dir(filename), "../templates/"+name+".json")
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, NonExistingTemplateError
	}

	return file, nil
}
