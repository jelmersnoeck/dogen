package template

import (
	"encoding/json"
	"fmt"
	"github.com/jelmersnoeck/noscito/services/easypdf/block"
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
	Layout         Layout
	TemplateBlocks []block.Block `json:"blocks"`
	Blocks         []block.BlockItem
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

func (t *Template) MatchData(data map[string]interface{}) {
	for _, template_block := range t.TemplateBlocks {
		template_block.Unmarshal()
		t.Blocks = append(t.Blocks, template_block.Item)
	}
}
