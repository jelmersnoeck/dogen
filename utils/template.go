package utils

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

func LoadTemplate(name string) []byte {
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
