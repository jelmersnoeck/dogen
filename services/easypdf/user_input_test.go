package easypdf_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/services/easypdf"
)

func TestParse(t *testing.T) {
	block := new(mocks.Block)
	pdf := new(easypdf.EasyPDF)

	input := mockMappedInterface(`{"test_id":{"url":"test"}}`)

	ui := new(easypdf.UserInput)
	ui.Block = block
	ui.InputId = "test_id"

	expected := mockMappedInterface(`{"url":"test"}`)
	block.On("Parse", pdf, expected).Return(true)

	ui.Parse(pdf, input)
}

func TestLoad(t *testing.T) {
	data := make(map[string]interface{})

	byt := []byte(`{
		"input_id": "test",
		"block": {
			"type": "image"
		}
	}`)

	json.Unmarshal(byt, &data)

	ui := easypdf.UserInput{}
	ui.Load(data)

	assert.Equal(t, ui.InputId, "test")
	assert.IsType(t, &easypdf.Image{}, ui.Block)
}

func mockMappedInterface(data string) (input map[string]interface{}) {
	json.Unmarshal([]byte(data), &input)
	return
}
