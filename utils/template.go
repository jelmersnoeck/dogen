package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var NonExistingTemplateError = errors.New("Template does not exist.")

func LoadTemplate(name string) ([]byte, error) {
	dir, pwdErr := os.Getwd()

	if pwdErr != nil {
		return nil, errors.New("Could not find current path.")
	}

	filepath := path.Join(dir + "/templates/" + name + ".json")
	fmt.Println(filepath)
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, NonExistingTemplateError
	}

	return file, nil
}
