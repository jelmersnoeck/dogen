package pdf

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type JsonTemplate struct {
	layout Layout
}

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

func (t *JsonTemplate) Layout() Layout {
	return t.layout
}
