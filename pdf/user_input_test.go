package pdf_test

import (
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

	block_attributes := map[string]interface{}{"type": "image"}
	block_data := map[string]interface{}{"block_data": block_attributes}

	input_attributes := map[string]interface{}{"url": "my-url"}
	input_data := map[string]interface{}{"test-key": input_attributes}

	template.On("LoadBlock", block_attributes, input_attributes).Return(true)

	block.Load(template, block_data, input_data)

	template.AssertExpectations(t)
}
