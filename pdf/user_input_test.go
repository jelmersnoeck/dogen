package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserInputBlockSuite struct {
	suite.Suite
}

func TestUserInputBlockSuite(t *testing.T) {
	suite.Run(t, new(UserInputBlockSuite))
}

func (s *UserInputBlockSuite) TestParse() {
	input_block := &mocks.Block{}
	user_input := &pdf.UserInput{"test", input_block}
	mpdf := &mocks.MPdf{}

	input_block.On("Parse", mpdf).Return(true)

	user_input.Parse(mpdf)

	input_block.AssertExpectations(s.T())
}

func (s *UserInputBlockSuite) TestLoad() {
	template := &mocks.Template{}
	block := &mocks.Block{}
	ui := &pdf.UserInput{}
	ui.InputId = "test-key"

	block_attributes := map[string]interface{}{"type": "image"}
	block_data := map[string]interface{}{"block_data": block_attributes}

	input_attributes := map[string]interface{}{"url": "my-url"}
	input_data := map[string]interface{}{"test-key": input_attributes}

	template.On("LoadBlock", block_attributes, input_attributes).Return(block)

	ui.Load(template, block_data, input_data)

	template.AssertExpectations(s.T())

	assert.NotNil(s.T(), ui.Block)
}
