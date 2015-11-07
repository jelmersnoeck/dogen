package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

var NonExistingTemplateError = errors.New("Template does not exist.")

func LoadTemplate(name string) ([]byte, error) {

	envPath := os.Getenv("APP_PATH")
	if envPath == "" {
		_, filename, _, ok := runtime.Caller(1)

		if !ok {
			return nil, errors.New("Could not find current path.")
		}
		envPath = path.Dir(filename) + "/.."
	}

	filepath := path.Join(envPath + "/templates/fixtures/" + name + ".json")
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, NonExistingTemplateError
	}

	return file, nil
}
