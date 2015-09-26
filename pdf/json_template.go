package pdf

import (
	"encoding/json"
	"sync"

	"github.com/mitchellh/mapstructure"
)

// JsonTemplate resembles a template that loads data from a JSON string.
type JsonTemplate struct {
	raw_block_data interface{}
	layout         Layout
	blocks         []Block
	errors         []error
	wg             *sync.WaitGroup
}

// NewJsonTemplate creates a new JsonTemplate and populates the layout and block
// values.
func NewJsonTemplate(template_data []byte) (Template, error) {
	t := &JsonTemplate{}
	t.wg = &sync.WaitGroup{}

	var data map[string]interface{}
	jsonErr := json.Unmarshal(template_data, &data) // TODO: error handling

	if jsonErr != nil {
		return nil, jsonErr
	}

	if loadErr := t.loadLayout(data["layout"]); loadErr != nil {
		return nil, loadErr
	}

	t.raw_block_data = data["blocks"]

	return t, nil
}

// AddError adds an error to the template so we can see what didn't parse well.
func (t *JsonTemplate) AddError(err error) {
	t.errors = append(t.errors, err)
}

// Errors return all the errors that have occured while making the template.
func (t *JsonTemplate) Errors() []error {
	return t.errors
}

// Layout returns the layout parsed from the JSON data into a Layout interface.
func (t *JsonTemplate) Layout() Layout {
	return t.layout
}

// WaitGroup returns the sync.WaitGroup that is used for this template.
func (t *JsonTemplate) WaitGroup() *sync.WaitGroup {
	return t.wg
}

// LoadBlocks will load the blocks that are stored in the raw_blocks data and
// mix it up with the user input that we give it.
func (t *JsonTemplate) LoadBlocks(user_input interface{}) {
	var blocks []map[string]interface{}
	var input map[string]interface{}
	mapstructure.Decode(t.raw_block_data, &blocks)
	mapstructure.Decode(user_input, &input)

	for _, block := range blocks {
		t.blocks = append(t.blocks, t.LoadBlock(block, input))
	}

	t.wg.Wait()
}

// LoadBlock takes a signle block item and a map of input data and loads the
// block and populates it with it's respective dataset.
func (t *JsonTemplate) LoadBlock(raw_block, raw_input map[string]interface{}) Block {
	block_type := raw_block["type"].(string)

	var block Block

	switch block_type {
	case "image":
		block = &Image{}
	case "user_input":
		block = &UserInput{}
	case "text_box":
		block = &TextBox{}
	case "rectangle":
		block = &Rectangle{}
	case "line":
		block = &Line{}
	case "add_page":
		block = &AddPage{}
	}

	mapstructure.Decode(raw_block, block)
	block.Load(t, raw_block, raw_input)

	return block
}

// Blocks returns the blocks parsed from the JSON data into an array of Blocks
func (t *JsonTemplate) Blocks() []Block {
	return t.blocks
}

func (t *JsonTemplate) loadLayout(data interface{}) error {
	var layout_data map[string]interface{}
	mapstructure.Decode(data, &layout_data)

	var err error
	t.layout, err = NewPageLayout(
		layout_data["orientation"].(string),
		layout_data["unit"].(string),
		layout_data["width"].(float64),
		layout_data["height"].(float64),
	)
	t.layout.SetFonts(loadFonts(layout_data))

	if err != nil {
		return err
	}

	return nil
}

func loadFonts(data map[string]interface{}) (fonts map[string]map[string]string) {
	if fontData, ok := data["fonts"]; ok {
		for fontName, fontStyles := range fontData.(map[string]interface{}) {
			fontItem := map[string]string{}

			for style, file := range fontStyles.(map[string]interface{}) {
				fontItem[style] = file.(string) + ".json"
			}

			fonts[fontName] = fontItem
		}
	}

	return
}
