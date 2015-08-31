package easypdf

import (
	"encoding/json"
	"fmt"
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
	Blocks []Block `json:"blocks"`
}

func LoadTemplate(name string) (t *Template) {
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
	t.Blocks = t.LoadBlocks(data["blocks"])

	return
}

func (t *Template) LoadBlocks(blocks interface{}) (b []Block) {
	b = make([]Block, len(t.Blocks))
	items := blocks.([]interface{})
	for _, item := range items {
		block := NewBlock(item)
		b = append(b, block)
	}
	return b
}
