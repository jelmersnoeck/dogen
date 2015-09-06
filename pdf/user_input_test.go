package pdf_test

import (
	"encoding/json"
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/pdf"
)

func TestParse(t *testing.T) {
	input_block := &mocks.Block{}
	user_input := &pdf.UserInput{"test", input_block}
	mpdf := &mocks.MPdf{}

	input_block.On("Parse", mpdf).Return(true)

	user_input.Parse(mpdf)

	input_block.AssertExpectations(t)
}

func TestLoad(t *testing.T) {
	template := &mocks.MTemplate{}
	block := &pdf.UserInput{}
	block.InputId = "test-key"

	block_data := map[string]interface{}{"key": ""}

	user_data := []byte(`{
		"test-key": {
			"type": "image"
		}
	}`)
	var user_input map[string]interface{}
	json.Unmarshal(user_data, &user_input)

	input_data := []byte(`{
		"type": "image"
	}`)
	var input_attributes map[string]interface{}
	json.Unmarshal(input_data, &input_attributes)

	template.On("LoadBlock", block_data, input_attributes).Return(true)

	block.Load(template, block_data, user_input)

	template.AssertExpectations(t)
}
