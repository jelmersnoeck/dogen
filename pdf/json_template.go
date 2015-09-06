package pdf

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// JsonTemplate resembles a template that loads data from a JSON string.
type JsonTemplate struct {
	layout Layout
	blocks []Block
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

	var layout_data map[string]interface{}
	mapstructure.Decode(data["layout"], &layout_data)

	t.layout, _ = NewPageLayout(
		layout_data["orientation"].(string),
		layout_data["unit"].(string),
		layout_data["width"].(float64),
		layout_data["height"].(float64),
	)

	return t, true
}

// Layout returns the layout parsed from the JSON data into a Layout interface.
func (t *JsonTemplate) Layout() Layout {
	return t.layout
}

// Blocks returns the blocks parsed from the JSON data into an array of Blocks
func (t *JsonTemplate) Blocks() []Block {
	return t.blocks
}
