package pdf

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// JsonTemplate resembles a template that loads data from a JSON string.
type JsonTemplate struct {
	raw_block_data interface{}
	layout         Layout
	blocks         []Block
}

// NewJsonTemplate creates a new JsonTemplate and populates the layout and block
// values.
func NewJsonTemplate(template_data []byte) (Template, bool) {
	t := &JsonTemplate{}

	var data map[string]interface{}
	err := json.Unmarshal(template_data, &data) // TODO: error handling

	if err != nil {
		return nil, false
	}

	t.loadLayout(data["layout"])
	t.raw_block_data = data["blocks"]

	return t, true
}

// Layout returns the layout parsed from the JSON data into a Layout interface.
func (t *JsonTemplate) Layout() Layout {
	return t.layout
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
	}

	mapstructure.Decode(raw_block, block)
	block.Load(t, raw_block, raw_input)

	return block
}

// Blocks returns the blocks parsed from the JSON data into an array of Blocks
func (t *JsonTemplate) Blocks() []Block {
	return t.blocks
}

func (t *JsonTemplate) loadLayout(data interface{}) {
	var layout_data map[string]interface{}
	mapstructure.Decode(data, &layout_data)

	t.layout, _ = NewPageLayout(
		layout_data["orientation"].(string),
		layout_data["unit"].(string),
		layout_data["width"].(float64),
		layout_data["height"].(float64),
	)
}
