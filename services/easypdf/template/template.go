package template

import (
	"encoding/json"
	"fmt"
	"github.com/jelmersnoeck/noscito/services/easypdf/block"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"path"
	"runtime"
)

type Layout struct {
	Width       float64
	Height      float64
	Unit        string
	Orientation string
}

type Template struct {
	Layout Layout
	Blocks []block.Block `json:"blocks"`
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

	var data map[string]interface{}
	json.Unmarshal(file, &data) // TODO: error handling

	mapstructure.Decode(data["layout"], &t.Layout)
	t.Blocks = block.LoadBlocks(data["blocks"], t.Blocks)

	return
}
