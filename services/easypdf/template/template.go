package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

type Block struct {
}

type Layout struct {
	Width  float64
	Height float64
	Unit   string
}

type Template struct {
	Layout Layout
	Blocks []Block
}

func Load(name string) (t *Template) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return
	}

	filepath := path.Join(path.Dir(filename), "../templates/"+name+".json")
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	t = new(Template)

	json.Unmarshal(file, &t)

	return
}
